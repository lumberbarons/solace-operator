# MsgVpnRestDeliveryPointRestConsumer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationHttpBasicUsername** | Pointer to **string** | The username that the REST Consumer will use to login to the REST host. | [optional] 
**AuthenticationHttpHeaderName** | Pointer to **string** | The authentication header name. Available since 2.15. | [optional] 
**AuthenticationOauthClientId** | Pointer to **string** | The OAuth client ID. Available since 2.19. | [optional] 
**AuthenticationOauthClientLastFailureReason** | Pointer to **string** | The reason for the most recent OAuth token retrieval failure. Available since 2.19. | [optional] 
**AuthenticationOauthClientLastFailureTime** | Pointer to **int32** | The time of the last OAuth token retrieval failure. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). Available since 2.19. | [optional] 
**AuthenticationOauthClientScope** | Pointer to **string** | The OAuth scope. Available since 2.19. | [optional] 
**AuthenticationOauthClientTokenEndpoint** | Pointer to **string** | The OAuth token endpoint URL that the REST Consumer will use to request a token for login to the REST host. Must begin with \&quot;https\&quot;. Available since 2.19. | [optional] 
**AuthenticationOauthClientTokenLifetime** | Pointer to **int64** | The validity duration of the OAuth token. Available since 2.19. | [optional] 
**AuthenticationOauthClientTokenRetrievedTime** | Pointer to **int32** | The time at which the broker requested the token from the OAuth token endpoint. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). Available since 2.19. | [optional] 
**AuthenticationOauthClientTokenState** | Pointer to **string** | The current state of the current OAuth token. The allowed values and their meaning are:  &lt;pre&gt; \&quot;valid\&quot; - The token is valid. \&quot;invalid\&quot; - The token is invalid. &lt;/pre&gt;  Available since 2.19. | [optional] 
**AuthenticationOauthJwtLastFailureReason** | Pointer to **string** | The reason for the most recent OAuth token retrieval failure. Available since 2.21. | [optional] 
**AuthenticationOauthJwtLastFailureTime** | Pointer to **int32** | The time of the last OAuth token retrieval failure. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). Available since 2.21. | [optional] 
**AuthenticationOauthJwtTokenEndpoint** | Pointer to **string** | The OAuth token endpoint URL that the REST Consumer will use to request a token for login to the REST host. Available since 2.21. | [optional] 
**AuthenticationOauthJwtTokenLifetime** | Pointer to **int64** | The validity duration of the OAuth token. Available since 2.21. | [optional] 
**AuthenticationOauthJwtTokenRetrievedTime** | Pointer to **int32** | The time at which the broker requested the token from the OAuth token endpoint. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). Available since 2.21. | [optional] 
**AuthenticationOauthJwtTokenState** | Pointer to **string** | The current state of the current OAuth token. The allowed values and their meaning are:  &lt;pre&gt; \&quot;valid\&quot; - The token is valid. \&quot;invalid\&quot; - The token is invalid. &lt;/pre&gt;  Available since 2.21. | [optional] 
**AuthenticationScheme** | Pointer to **string** | The authentication scheme used by the REST Consumer to login to the REST host. The allowed values and their meaning are:  &lt;pre&gt; \&quot;none\&quot; - Login with no authentication. This may be useful for anonymous connections or when a REST Consumer does not require authentication. \&quot;http-basic\&quot; - Login with a username and optional password according to HTTP Basic authentication as per RFC2616. \&quot;client-certificate\&quot; - Login with a client TLS certificate as per RFC5246. Client certificate authentication is only available on TLS connections. \&quot;http-header\&quot; - Login with a specified HTTP header. \&quot;oauth-client\&quot; - Login with OAuth 2.0 client credentials. \&quot;oauth-jwt\&quot; - Login with OAuth (RFC 7523 JWT Profile). \&quot;transparent\&quot; - Login using the Authorization header from the message properties, if present. Transparent authentication passes along existing Authorization header metadata instead of discarding it. Note that if the message is coming from a REST producer, the REST service must be configured to forward the Authorization header. &lt;/pre&gt;  | [optional] 
**Counter** | Pointer to [**MsgVpnRestDeliveryPointRestConsumerCounter**](MsgVpnRestDeliveryPointRestConsumerCounter.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the REST Consumer is enabled. | [optional] 
**HttpMethod** | Pointer to **string** | The HTTP method to use (POST or PUT). This is used only when operating in the REST service \&quot;messaging\&quot; mode and is ignored in \&quot;gateway\&quot; mode. The allowed values and their meaning are:  &lt;pre&gt; \&quot;post\&quot; - Use the POST HTTP method. \&quot;put\&quot; - Use the PUT HTTP method. &lt;/pre&gt;  Available since 2.17. | [optional] 
**HttpRequestConnectionCloseTxMsgCount** | Pointer to **int64** | The number of HTTP request messages transmitted to the REST Consumer to close the connection. Available since 2.13. | [optional] 
**HttpRequestOutstandingTxMsgCount** | Pointer to **int64** | The number of HTTP request messages transmitted to the REST Consumer that are waiting for a response. Available since 2.13. | [optional] 
**HttpRequestTimedOutTxMsgCount** | Pointer to **int64** | The number of HTTP request messages transmitted to the REST Consumer that have timed out. Available since 2.13. | [optional] 
**HttpRequestTxByteCount** | Pointer to **int64** | The amount of HTTP request messages transmitted to the REST Consumer, in bytes (B). Available since 2.13. | [optional] 
**HttpRequestTxMsgCount** | Pointer to **int64** | The number of HTTP request messages transmitted to the REST Consumer. Available since 2.13. | [optional] 
**HttpResponseErrorRxMsgCount** | Pointer to **int64** | The number of HTTP client/server error response messages received from the REST Consumer. Available since 2.13. | [optional] 
**HttpResponseRxByteCount** | Pointer to **int64** | The amount of HTTP response messages received from the REST Consumer, in bytes (B). Available since 2.13. | [optional] 
**HttpResponseRxMsgCount** | Pointer to **int64** | The number of HTTP response messages received from the REST Consumer. Available since 2.13. | [optional] 
**HttpResponseSuccessRxMsgCount** | Pointer to **int64** | The number of HTTP successful response messages received from the REST Consumer. Available since 2.13. | [optional] 
**LastConnectionFailureLocalEndpoint** | Pointer to **string** | The local endpoint at the time of the last connection failure. | [optional] 
**LastConnectionFailureReason** | Pointer to **string** | The reason for the last connection failure between local and remote endpoints. | [optional] 
**LastConnectionFailureRemoteEndpoint** | Pointer to **string** | The remote endpoint at the time of the last connection failure. | [optional] 
**LastConnectionFailureTime** | Pointer to **int32** | The timestamp of the last connection failure between local and remote endpoints. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] 
**LastFailureReason** | Pointer to **string** | The reason for the last REST Consumer failure. | [optional] 
**LastFailureTime** | Pointer to **int32** | The timestamp of the last REST Consumer failure. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] 
**LocalInterface** | Pointer to **string** | The interface that will be used for all outgoing connections associated with the REST Consumer. When unspecified, an interface is automatically chosen. | [optional] 
**MaxPostWaitTime** | Pointer to **int32** | The maximum amount of time (in seconds) to wait for an HTTP POST response from the REST Consumer. Once this time is exceeded, the TCP connection is reset. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**OutgoingConnectionCount** | Pointer to **int32** | The number of concurrent TCP connections open to the REST Consumer. | [optional] 
**RemoteHost** | Pointer to **string** | The IP address or DNS name for the REST Consumer. | [optional] 
**RemoteOutgoingConnectionUpCount** | Pointer to **int64** | The number of outgoing connections for the REST Consumer that are up. | [optional] 
**RemotePort** | Pointer to **int64** | The port associated with the host of the REST Consumer. | [optional] 
**RestConsumerName** | Pointer to **string** | The name of the REST Consumer. | [optional] 
**RestDeliveryPointName** | Pointer to **string** | The name of the REST Delivery Point. | [optional] 
**RetryDelay** | Pointer to **int32** | The number of seconds that must pass before retrying the remote REST Consumer connection. | [optional] 
**TlsCipherSuiteList** | Pointer to **string** | The colon-separated list of cipher suites the REST Consumer uses in its encrypted connection. The value &#x60;\&quot;default\&quot;&#x60; implies all supported suites ordered from most secure to least secure. The list of default cipher suites is available in the &#x60;tlsCipherSuiteMsgBackboneDefaultList&#x60; attribute of the Broker object in the Monitoring API. The REST Consumer should choose the first suite from this list that it supports. | [optional] 
**TlsEnabled** | Pointer to **bool** | Indicates whether encryption (TLS) is enabled for the REST Consumer. | [optional] 
**Up** | Pointer to **bool** | Indicates whether the operational state of the REST Consumer is up. | [optional] 

## Methods

### NewMsgVpnRestDeliveryPointRestConsumer

`func NewMsgVpnRestDeliveryPointRestConsumer() *MsgVpnRestDeliveryPointRestConsumer`

NewMsgVpnRestDeliveryPointRestConsumer instantiates a new MsgVpnRestDeliveryPointRestConsumer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnRestDeliveryPointRestConsumerWithDefaults

`func NewMsgVpnRestDeliveryPointRestConsumerWithDefaults() *MsgVpnRestDeliveryPointRestConsumer`

NewMsgVpnRestDeliveryPointRestConsumerWithDefaults instantiates a new MsgVpnRestDeliveryPointRestConsumer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationHttpBasicUsername

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationHttpBasicUsername() string`

GetAuthenticationHttpBasicUsername returns the AuthenticationHttpBasicUsername field if non-nil, zero value otherwise.

### GetAuthenticationHttpBasicUsernameOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationHttpBasicUsernameOk() (*string, bool)`

GetAuthenticationHttpBasicUsernameOk returns a tuple with the AuthenticationHttpBasicUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationHttpBasicUsername

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationHttpBasicUsername(v string)`

SetAuthenticationHttpBasicUsername sets AuthenticationHttpBasicUsername field to given value.

### HasAuthenticationHttpBasicUsername

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationHttpBasicUsername() bool`

HasAuthenticationHttpBasicUsername returns a boolean if a field has been set.

### GetAuthenticationHttpHeaderName

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationHttpHeaderName() string`

GetAuthenticationHttpHeaderName returns the AuthenticationHttpHeaderName field if non-nil, zero value otherwise.

### GetAuthenticationHttpHeaderNameOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationHttpHeaderNameOk() (*string, bool)`

GetAuthenticationHttpHeaderNameOk returns a tuple with the AuthenticationHttpHeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationHttpHeaderName

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationHttpHeaderName(v string)`

SetAuthenticationHttpHeaderName sets AuthenticationHttpHeaderName field to given value.

### HasAuthenticationHttpHeaderName

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationHttpHeaderName() bool`

HasAuthenticationHttpHeaderName returns a boolean if a field has been set.

### GetAuthenticationOauthClientId

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthClientId() string`

GetAuthenticationOauthClientId returns the AuthenticationOauthClientId field if non-nil, zero value otherwise.

### GetAuthenticationOauthClientIdOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthClientIdOk() (*string, bool)`

GetAuthenticationOauthClientIdOk returns a tuple with the AuthenticationOauthClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationOauthClientId

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationOauthClientId(v string)`

SetAuthenticationOauthClientId sets AuthenticationOauthClientId field to given value.

### HasAuthenticationOauthClientId

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationOauthClientId() bool`

HasAuthenticationOauthClientId returns a boolean if a field has been set.

### GetAuthenticationOauthClientLastFailureReason

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthClientLastFailureReason() string`

GetAuthenticationOauthClientLastFailureReason returns the AuthenticationOauthClientLastFailureReason field if non-nil, zero value otherwise.

### GetAuthenticationOauthClientLastFailureReasonOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthClientLastFailureReasonOk() (*string, bool)`

GetAuthenticationOauthClientLastFailureReasonOk returns a tuple with the AuthenticationOauthClientLastFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationOauthClientLastFailureReason

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationOauthClientLastFailureReason(v string)`

SetAuthenticationOauthClientLastFailureReason sets AuthenticationOauthClientLastFailureReason field to given value.

### HasAuthenticationOauthClientLastFailureReason

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationOauthClientLastFailureReason() bool`

HasAuthenticationOauthClientLastFailureReason returns a boolean if a field has been set.

### GetAuthenticationOauthClientLastFailureTime

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthClientLastFailureTime() int32`

GetAuthenticationOauthClientLastFailureTime returns the AuthenticationOauthClientLastFailureTime field if non-nil, zero value otherwise.

### GetAuthenticationOauthClientLastFailureTimeOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthClientLastFailureTimeOk() (*int32, bool)`

GetAuthenticationOauthClientLastFailureTimeOk returns a tuple with the AuthenticationOauthClientLastFailureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationOauthClientLastFailureTime

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationOauthClientLastFailureTime(v int32)`

SetAuthenticationOauthClientLastFailureTime sets AuthenticationOauthClientLastFailureTime field to given value.

### HasAuthenticationOauthClientLastFailureTime

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationOauthClientLastFailureTime() bool`

HasAuthenticationOauthClientLastFailureTime returns a boolean if a field has been set.

### GetAuthenticationOauthClientScope

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthClientScope() string`

GetAuthenticationOauthClientScope returns the AuthenticationOauthClientScope field if non-nil, zero value otherwise.

### GetAuthenticationOauthClientScopeOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthClientScopeOk() (*string, bool)`

GetAuthenticationOauthClientScopeOk returns a tuple with the AuthenticationOauthClientScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationOauthClientScope

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationOauthClientScope(v string)`

SetAuthenticationOauthClientScope sets AuthenticationOauthClientScope field to given value.

### HasAuthenticationOauthClientScope

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationOauthClientScope() bool`

HasAuthenticationOauthClientScope returns a boolean if a field has been set.

### GetAuthenticationOauthClientTokenEndpoint

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthClientTokenEndpoint() string`

GetAuthenticationOauthClientTokenEndpoint returns the AuthenticationOauthClientTokenEndpoint field if non-nil, zero value otherwise.

### GetAuthenticationOauthClientTokenEndpointOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthClientTokenEndpointOk() (*string, bool)`

GetAuthenticationOauthClientTokenEndpointOk returns a tuple with the AuthenticationOauthClientTokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationOauthClientTokenEndpoint

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationOauthClientTokenEndpoint(v string)`

SetAuthenticationOauthClientTokenEndpoint sets AuthenticationOauthClientTokenEndpoint field to given value.

### HasAuthenticationOauthClientTokenEndpoint

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationOauthClientTokenEndpoint() bool`

HasAuthenticationOauthClientTokenEndpoint returns a boolean if a field has been set.

### GetAuthenticationOauthClientTokenLifetime

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthClientTokenLifetime() int64`

GetAuthenticationOauthClientTokenLifetime returns the AuthenticationOauthClientTokenLifetime field if non-nil, zero value otherwise.

### GetAuthenticationOauthClientTokenLifetimeOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthClientTokenLifetimeOk() (*int64, bool)`

GetAuthenticationOauthClientTokenLifetimeOk returns a tuple with the AuthenticationOauthClientTokenLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationOauthClientTokenLifetime

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationOauthClientTokenLifetime(v int64)`

SetAuthenticationOauthClientTokenLifetime sets AuthenticationOauthClientTokenLifetime field to given value.

### HasAuthenticationOauthClientTokenLifetime

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationOauthClientTokenLifetime() bool`

HasAuthenticationOauthClientTokenLifetime returns a boolean if a field has been set.

### GetAuthenticationOauthClientTokenRetrievedTime

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthClientTokenRetrievedTime() int32`

GetAuthenticationOauthClientTokenRetrievedTime returns the AuthenticationOauthClientTokenRetrievedTime field if non-nil, zero value otherwise.

### GetAuthenticationOauthClientTokenRetrievedTimeOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthClientTokenRetrievedTimeOk() (*int32, bool)`

GetAuthenticationOauthClientTokenRetrievedTimeOk returns a tuple with the AuthenticationOauthClientTokenRetrievedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationOauthClientTokenRetrievedTime

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationOauthClientTokenRetrievedTime(v int32)`

SetAuthenticationOauthClientTokenRetrievedTime sets AuthenticationOauthClientTokenRetrievedTime field to given value.

### HasAuthenticationOauthClientTokenRetrievedTime

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationOauthClientTokenRetrievedTime() bool`

HasAuthenticationOauthClientTokenRetrievedTime returns a boolean if a field has been set.

### GetAuthenticationOauthClientTokenState

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthClientTokenState() string`

GetAuthenticationOauthClientTokenState returns the AuthenticationOauthClientTokenState field if non-nil, zero value otherwise.

### GetAuthenticationOauthClientTokenStateOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthClientTokenStateOk() (*string, bool)`

GetAuthenticationOauthClientTokenStateOk returns a tuple with the AuthenticationOauthClientTokenState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationOauthClientTokenState

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationOauthClientTokenState(v string)`

SetAuthenticationOauthClientTokenState sets AuthenticationOauthClientTokenState field to given value.

### HasAuthenticationOauthClientTokenState

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationOauthClientTokenState() bool`

HasAuthenticationOauthClientTokenState returns a boolean if a field has been set.

### GetAuthenticationOauthJwtLastFailureReason

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthJwtLastFailureReason() string`

GetAuthenticationOauthJwtLastFailureReason returns the AuthenticationOauthJwtLastFailureReason field if non-nil, zero value otherwise.

### GetAuthenticationOauthJwtLastFailureReasonOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthJwtLastFailureReasonOk() (*string, bool)`

GetAuthenticationOauthJwtLastFailureReasonOk returns a tuple with the AuthenticationOauthJwtLastFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationOauthJwtLastFailureReason

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationOauthJwtLastFailureReason(v string)`

SetAuthenticationOauthJwtLastFailureReason sets AuthenticationOauthJwtLastFailureReason field to given value.

### HasAuthenticationOauthJwtLastFailureReason

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationOauthJwtLastFailureReason() bool`

HasAuthenticationOauthJwtLastFailureReason returns a boolean if a field has been set.

### GetAuthenticationOauthJwtLastFailureTime

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthJwtLastFailureTime() int32`

GetAuthenticationOauthJwtLastFailureTime returns the AuthenticationOauthJwtLastFailureTime field if non-nil, zero value otherwise.

### GetAuthenticationOauthJwtLastFailureTimeOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthJwtLastFailureTimeOk() (*int32, bool)`

GetAuthenticationOauthJwtLastFailureTimeOk returns a tuple with the AuthenticationOauthJwtLastFailureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationOauthJwtLastFailureTime

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationOauthJwtLastFailureTime(v int32)`

SetAuthenticationOauthJwtLastFailureTime sets AuthenticationOauthJwtLastFailureTime field to given value.

### HasAuthenticationOauthJwtLastFailureTime

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationOauthJwtLastFailureTime() bool`

HasAuthenticationOauthJwtLastFailureTime returns a boolean if a field has been set.

### GetAuthenticationOauthJwtTokenEndpoint

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthJwtTokenEndpoint() string`

GetAuthenticationOauthJwtTokenEndpoint returns the AuthenticationOauthJwtTokenEndpoint field if non-nil, zero value otherwise.

### GetAuthenticationOauthJwtTokenEndpointOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthJwtTokenEndpointOk() (*string, bool)`

GetAuthenticationOauthJwtTokenEndpointOk returns a tuple with the AuthenticationOauthJwtTokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationOauthJwtTokenEndpoint

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationOauthJwtTokenEndpoint(v string)`

SetAuthenticationOauthJwtTokenEndpoint sets AuthenticationOauthJwtTokenEndpoint field to given value.

### HasAuthenticationOauthJwtTokenEndpoint

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationOauthJwtTokenEndpoint() bool`

HasAuthenticationOauthJwtTokenEndpoint returns a boolean if a field has been set.

### GetAuthenticationOauthJwtTokenLifetime

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthJwtTokenLifetime() int64`

GetAuthenticationOauthJwtTokenLifetime returns the AuthenticationOauthJwtTokenLifetime field if non-nil, zero value otherwise.

### GetAuthenticationOauthJwtTokenLifetimeOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthJwtTokenLifetimeOk() (*int64, bool)`

GetAuthenticationOauthJwtTokenLifetimeOk returns a tuple with the AuthenticationOauthJwtTokenLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationOauthJwtTokenLifetime

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationOauthJwtTokenLifetime(v int64)`

SetAuthenticationOauthJwtTokenLifetime sets AuthenticationOauthJwtTokenLifetime field to given value.

### HasAuthenticationOauthJwtTokenLifetime

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationOauthJwtTokenLifetime() bool`

HasAuthenticationOauthJwtTokenLifetime returns a boolean if a field has been set.

### GetAuthenticationOauthJwtTokenRetrievedTime

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthJwtTokenRetrievedTime() int32`

GetAuthenticationOauthJwtTokenRetrievedTime returns the AuthenticationOauthJwtTokenRetrievedTime field if non-nil, zero value otherwise.

### GetAuthenticationOauthJwtTokenRetrievedTimeOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthJwtTokenRetrievedTimeOk() (*int32, bool)`

GetAuthenticationOauthJwtTokenRetrievedTimeOk returns a tuple with the AuthenticationOauthJwtTokenRetrievedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationOauthJwtTokenRetrievedTime

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationOauthJwtTokenRetrievedTime(v int32)`

SetAuthenticationOauthJwtTokenRetrievedTime sets AuthenticationOauthJwtTokenRetrievedTime field to given value.

### HasAuthenticationOauthJwtTokenRetrievedTime

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationOauthJwtTokenRetrievedTime() bool`

HasAuthenticationOauthJwtTokenRetrievedTime returns a boolean if a field has been set.

### GetAuthenticationOauthJwtTokenState

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthJwtTokenState() string`

GetAuthenticationOauthJwtTokenState returns the AuthenticationOauthJwtTokenState field if non-nil, zero value otherwise.

### GetAuthenticationOauthJwtTokenStateOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthJwtTokenStateOk() (*string, bool)`

GetAuthenticationOauthJwtTokenStateOk returns a tuple with the AuthenticationOauthJwtTokenState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationOauthJwtTokenState

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationOauthJwtTokenState(v string)`

SetAuthenticationOauthJwtTokenState sets AuthenticationOauthJwtTokenState field to given value.

### HasAuthenticationOauthJwtTokenState

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationOauthJwtTokenState() bool`

HasAuthenticationOauthJwtTokenState returns a boolean if a field has been set.

### GetAuthenticationScheme

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationScheme() string`

GetAuthenticationScheme returns the AuthenticationScheme field if non-nil, zero value otherwise.

### GetAuthenticationSchemeOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationSchemeOk() (*string, bool)`

GetAuthenticationSchemeOk returns a tuple with the AuthenticationScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationScheme

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationScheme(v string)`

SetAuthenticationScheme sets AuthenticationScheme field to given value.

### HasAuthenticationScheme

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationScheme() bool`

HasAuthenticationScheme returns a boolean if a field has been set.

### GetCounter

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetCounter() MsgVpnRestDeliveryPointRestConsumerCounter`

GetCounter returns the Counter field if non-nil, zero value otherwise.

### GetCounterOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetCounterOk() (*MsgVpnRestDeliveryPointRestConsumerCounter, bool)`

GetCounterOk returns a tuple with the Counter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounter

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetCounter(v MsgVpnRestDeliveryPointRestConsumerCounter)`

SetCounter sets Counter field to given value.

### HasCounter

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasCounter() bool`

HasCounter returns a boolean if a field has been set.

### GetEnabled

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHttpMethod

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpMethod() string`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpMethodOk() (*string, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetHttpMethod(v string)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### GetHttpRequestConnectionCloseTxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpRequestConnectionCloseTxMsgCount() int64`

GetHttpRequestConnectionCloseTxMsgCount returns the HttpRequestConnectionCloseTxMsgCount field if non-nil, zero value otherwise.

### GetHttpRequestConnectionCloseTxMsgCountOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpRequestConnectionCloseTxMsgCountOk() (*int64, bool)`

GetHttpRequestConnectionCloseTxMsgCountOk returns a tuple with the HttpRequestConnectionCloseTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpRequestConnectionCloseTxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetHttpRequestConnectionCloseTxMsgCount(v int64)`

SetHttpRequestConnectionCloseTxMsgCount sets HttpRequestConnectionCloseTxMsgCount field to given value.

### HasHttpRequestConnectionCloseTxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasHttpRequestConnectionCloseTxMsgCount() bool`

HasHttpRequestConnectionCloseTxMsgCount returns a boolean if a field has been set.

### GetHttpRequestOutstandingTxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpRequestOutstandingTxMsgCount() int64`

GetHttpRequestOutstandingTxMsgCount returns the HttpRequestOutstandingTxMsgCount field if non-nil, zero value otherwise.

### GetHttpRequestOutstandingTxMsgCountOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpRequestOutstandingTxMsgCountOk() (*int64, bool)`

GetHttpRequestOutstandingTxMsgCountOk returns a tuple with the HttpRequestOutstandingTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpRequestOutstandingTxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetHttpRequestOutstandingTxMsgCount(v int64)`

SetHttpRequestOutstandingTxMsgCount sets HttpRequestOutstandingTxMsgCount field to given value.

### HasHttpRequestOutstandingTxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasHttpRequestOutstandingTxMsgCount() bool`

HasHttpRequestOutstandingTxMsgCount returns a boolean if a field has been set.

### GetHttpRequestTimedOutTxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpRequestTimedOutTxMsgCount() int64`

GetHttpRequestTimedOutTxMsgCount returns the HttpRequestTimedOutTxMsgCount field if non-nil, zero value otherwise.

### GetHttpRequestTimedOutTxMsgCountOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpRequestTimedOutTxMsgCountOk() (*int64, bool)`

GetHttpRequestTimedOutTxMsgCountOk returns a tuple with the HttpRequestTimedOutTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpRequestTimedOutTxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetHttpRequestTimedOutTxMsgCount(v int64)`

SetHttpRequestTimedOutTxMsgCount sets HttpRequestTimedOutTxMsgCount field to given value.

### HasHttpRequestTimedOutTxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasHttpRequestTimedOutTxMsgCount() bool`

HasHttpRequestTimedOutTxMsgCount returns a boolean if a field has been set.

### GetHttpRequestTxByteCount

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpRequestTxByteCount() int64`

GetHttpRequestTxByteCount returns the HttpRequestTxByteCount field if non-nil, zero value otherwise.

### GetHttpRequestTxByteCountOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpRequestTxByteCountOk() (*int64, bool)`

GetHttpRequestTxByteCountOk returns a tuple with the HttpRequestTxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpRequestTxByteCount

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetHttpRequestTxByteCount(v int64)`

SetHttpRequestTxByteCount sets HttpRequestTxByteCount field to given value.

### HasHttpRequestTxByteCount

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasHttpRequestTxByteCount() bool`

HasHttpRequestTxByteCount returns a boolean if a field has been set.

### GetHttpRequestTxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpRequestTxMsgCount() int64`

GetHttpRequestTxMsgCount returns the HttpRequestTxMsgCount field if non-nil, zero value otherwise.

### GetHttpRequestTxMsgCountOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpRequestTxMsgCountOk() (*int64, bool)`

GetHttpRequestTxMsgCountOk returns a tuple with the HttpRequestTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpRequestTxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetHttpRequestTxMsgCount(v int64)`

SetHttpRequestTxMsgCount sets HttpRequestTxMsgCount field to given value.

### HasHttpRequestTxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasHttpRequestTxMsgCount() bool`

HasHttpRequestTxMsgCount returns a boolean if a field has been set.

### GetHttpResponseErrorRxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpResponseErrorRxMsgCount() int64`

GetHttpResponseErrorRxMsgCount returns the HttpResponseErrorRxMsgCount field if non-nil, zero value otherwise.

### GetHttpResponseErrorRxMsgCountOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpResponseErrorRxMsgCountOk() (*int64, bool)`

GetHttpResponseErrorRxMsgCountOk returns a tuple with the HttpResponseErrorRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpResponseErrorRxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetHttpResponseErrorRxMsgCount(v int64)`

SetHttpResponseErrorRxMsgCount sets HttpResponseErrorRxMsgCount field to given value.

### HasHttpResponseErrorRxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasHttpResponseErrorRxMsgCount() bool`

HasHttpResponseErrorRxMsgCount returns a boolean if a field has been set.

### GetHttpResponseRxByteCount

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpResponseRxByteCount() int64`

GetHttpResponseRxByteCount returns the HttpResponseRxByteCount field if non-nil, zero value otherwise.

### GetHttpResponseRxByteCountOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpResponseRxByteCountOk() (*int64, bool)`

GetHttpResponseRxByteCountOk returns a tuple with the HttpResponseRxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpResponseRxByteCount

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetHttpResponseRxByteCount(v int64)`

SetHttpResponseRxByteCount sets HttpResponseRxByteCount field to given value.

### HasHttpResponseRxByteCount

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasHttpResponseRxByteCount() bool`

HasHttpResponseRxByteCount returns a boolean if a field has been set.

### GetHttpResponseRxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpResponseRxMsgCount() int64`

GetHttpResponseRxMsgCount returns the HttpResponseRxMsgCount field if non-nil, zero value otherwise.

### GetHttpResponseRxMsgCountOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpResponseRxMsgCountOk() (*int64, bool)`

GetHttpResponseRxMsgCountOk returns a tuple with the HttpResponseRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpResponseRxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetHttpResponseRxMsgCount(v int64)`

SetHttpResponseRxMsgCount sets HttpResponseRxMsgCount field to given value.

### HasHttpResponseRxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasHttpResponseRxMsgCount() bool`

HasHttpResponseRxMsgCount returns a boolean if a field has been set.

### GetHttpResponseSuccessRxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpResponseSuccessRxMsgCount() int64`

GetHttpResponseSuccessRxMsgCount returns the HttpResponseSuccessRxMsgCount field if non-nil, zero value otherwise.

### GetHttpResponseSuccessRxMsgCountOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetHttpResponseSuccessRxMsgCountOk() (*int64, bool)`

GetHttpResponseSuccessRxMsgCountOk returns a tuple with the HttpResponseSuccessRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpResponseSuccessRxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetHttpResponseSuccessRxMsgCount(v int64)`

SetHttpResponseSuccessRxMsgCount sets HttpResponseSuccessRxMsgCount field to given value.

### HasHttpResponseSuccessRxMsgCount

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasHttpResponseSuccessRxMsgCount() bool`

HasHttpResponseSuccessRxMsgCount returns a boolean if a field has been set.

### GetLastConnectionFailureLocalEndpoint

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetLastConnectionFailureLocalEndpoint() string`

GetLastConnectionFailureLocalEndpoint returns the LastConnectionFailureLocalEndpoint field if non-nil, zero value otherwise.

### GetLastConnectionFailureLocalEndpointOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetLastConnectionFailureLocalEndpointOk() (*string, bool)`

GetLastConnectionFailureLocalEndpointOk returns a tuple with the LastConnectionFailureLocalEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConnectionFailureLocalEndpoint

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetLastConnectionFailureLocalEndpoint(v string)`

SetLastConnectionFailureLocalEndpoint sets LastConnectionFailureLocalEndpoint field to given value.

### HasLastConnectionFailureLocalEndpoint

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasLastConnectionFailureLocalEndpoint() bool`

HasLastConnectionFailureLocalEndpoint returns a boolean if a field has been set.

### GetLastConnectionFailureReason

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetLastConnectionFailureReason() string`

GetLastConnectionFailureReason returns the LastConnectionFailureReason field if non-nil, zero value otherwise.

### GetLastConnectionFailureReasonOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetLastConnectionFailureReasonOk() (*string, bool)`

GetLastConnectionFailureReasonOk returns a tuple with the LastConnectionFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConnectionFailureReason

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetLastConnectionFailureReason(v string)`

SetLastConnectionFailureReason sets LastConnectionFailureReason field to given value.

### HasLastConnectionFailureReason

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasLastConnectionFailureReason() bool`

HasLastConnectionFailureReason returns a boolean if a field has been set.

### GetLastConnectionFailureRemoteEndpoint

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetLastConnectionFailureRemoteEndpoint() string`

GetLastConnectionFailureRemoteEndpoint returns the LastConnectionFailureRemoteEndpoint field if non-nil, zero value otherwise.

### GetLastConnectionFailureRemoteEndpointOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetLastConnectionFailureRemoteEndpointOk() (*string, bool)`

GetLastConnectionFailureRemoteEndpointOk returns a tuple with the LastConnectionFailureRemoteEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConnectionFailureRemoteEndpoint

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetLastConnectionFailureRemoteEndpoint(v string)`

SetLastConnectionFailureRemoteEndpoint sets LastConnectionFailureRemoteEndpoint field to given value.

### HasLastConnectionFailureRemoteEndpoint

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasLastConnectionFailureRemoteEndpoint() bool`

HasLastConnectionFailureRemoteEndpoint returns a boolean if a field has been set.

### GetLastConnectionFailureTime

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetLastConnectionFailureTime() int32`

GetLastConnectionFailureTime returns the LastConnectionFailureTime field if non-nil, zero value otherwise.

### GetLastConnectionFailureTimeOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetLastConnectionFailureTimeOk() (*int32, bool)`

GetLastConnectionFailureTimeOk returns a tuple with the LastConnectionFailureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConnectionFailureTime

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetLastConnectionFailureTime(v int32)`

SetLastConnectionFailureTime sets LastConnectionFailureTime field to given value.

### HasLastConnectionFailureTime

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasLastConnectionFailureTime() bool`

HasLastConnectionFailureTime returns a boolean if a field has been set.

### GetLastFailureReason

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetLastFailureReason() string`

GetLastFailureReason returns the LastFailureReason field if non-nil, zero value otherwise.

### GetLastFailureReasonOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetLastFailureReasonOk() (*string, bool)`

GetLastFailureReasonOk returns a tuple with the LastFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailureReason

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetLastFailureReason(v string)`

SetLastFailureReason sets LastFailureReason field to given value.

### HasLastFailureReason

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasLastFailureReason() bool`

HasLastFailureReason returns a boolean if a field has been set.

### GetLastFailureTime

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetLastFailureTime() int32`

GetLastFailureTime returns the LastFailureTime field if non-nil, zero value otherwise.

### GetLastFailureTimeOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetLastFailureTimeOk() (*int32, bool)`

GetLastFailureTimeOk returns a tuple with the LastFailureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailureTime

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetLastFailureTime(v int32)`

SetLastFailureTime sets LastFailureTime field to given value.

### HasLastFailureTime

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasLastFailureTime() bool`

HasLastFailureTime returns a boolean if a field has been set.

### GetLocalInterface

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetLocalInterface() string`

GetLocalInterface returns the LocalInterface field if non-nil, zero value otherwise.

### GetLocalInterfaceOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetLocalInterfaceOk() (*string, bool)`

GetLocalInterfaceOk returns a tuple with the LocalInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalInterface

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetLocalInterface(v string)`

SetLocalInterface sets LocalInterface field to given value.

### HasLocalInterface

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasLocalInterface() bool`

HasLocalInterface returns a boolean if a field has been set.

### GetMaxPostWaitTime

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetMaxPostWaitTime() int32`

GetMaxPostWaitTime returns the MaxPostWaitTime field if non-nil, zero value otherwise.

### GetMaxPostWaitTimeOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetMaxPostWaitTimeOk() (*int32, bool)`

GetMaxPostWaitTimeOk returns a tuple with the MaxPostWaitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPostWaitTime

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetMaxPostWaitTime(v int32)`

SetMaxPostWaitTime sets MaxPostWaitTime field to given value.

### HasMaxPostWaitTime

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasMaxPostWaitTime() bool`

HasMaxPostWaitTime returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetOutgoingConnectionCount

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetOutgoingConnectionCount() int32`

GetOutgoingConnectionCount returns the OutgoingConnectionCount field if non-nil, zero value otherwise.

### GetOutgoingConnectionCountOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetOutgoingConnectionCountOk() (*int32, bool)`

GetOutgoingConnectionCountOk returns a tuple with the OutgoingConnectionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingConnectionCount

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetOutgoingConnectionCount(v int32)`

SetOutgoingConnectionCount sets OutgoingConnectionCount field to given value.

### HasOutgoingConnectionCount

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasOutgoingConnectionCount() bool`

HasOutgoingConnectionCount returns a boolean if a field has been set.

### GetRemoteHost

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetRemoteHost() string`

GetRemoteHost returns the RemoteHost field if non-nil, zero value otherwise.

### GetRemoteHostOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetRemoteHostOk() (*string, bool)`

GetRemoteHostOk returns a tuple with the RemoteHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteHost

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetRemoteHost(v string)`

SetRemoteHost sets RemoteHost field to given value.

### HasRemoteHost

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasRemoteHost() bool`

HasRemoteHost returns a boolean if a field has been set.

### GetRemoteOutgoingConnectionUpCount

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetRemoteOutgoingConnectionUpCount() int64`

GetRemoteOutgoingConnectionUpCount returns the RemoteOutgoingConnectionUpCount field if non-nil, zero value otherwise.

### GetRemoteOutgoingConnectionUpCountOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetRemoteOutgoingConnectionUpCountOk() (*int64, bool)`

GetRemoteOutgoingConnectionUpCountOk returns a tuple with the RemoteOutgoingConnectionUpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteOutgoingConnectionUpCount

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetRemoteOutgoingConnectionUpCount(v int64)`

SetRemoteOutgoingConnectionUpCount sets RemoteOutgoingConnectionUpCount field to given value.

### HasRemoteOutgoingConnectionUpCount

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasRemoteOutgoingConnectionUpCount() bool`

HasRemoteOutgoingConnectionUpCount returns a boolean if a field has been set.

### GetRemotePort

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetRemotePort() int64`

GetRemotePort returns the RemotePort field if non-nil, zero value otherwise.

### GetRemotePortOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetRemotePortOk() (*int64, bool)`

GetRemotePortOk returns a tuple with the RemotePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePort

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetRemotePort(v int64)`

SetRemotePort sets RemotePort field to given value.

### HasRemotePort

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasRemotePort() bool`

HasRemotePort returns a boolean if a field has been set.

### GetRestConsumerName

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetRestConsumerName() string`

GetRestConsumerName returns the RestConsumerName field if non-nil, zero value otherwise.

### GetRestConsumerNameOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetRestConsumerNameOk() (*string, bool)`

GetRestConsumerNameOk returns a tuple with the RestConsumerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestConsumerName

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetRestConsumerName(v string)`

SetRestConsumerName sets RestConsumerName field to given value.

### HasRestConsumerName

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasRestConsumerName() bool`

HasRestConsumerName returns a boolean if a field has been set.

### GetRestDeliveryPointName

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetRestDeliveryPointName() string`

GetRestDeliveryPointName returns the RestDeliveryPointName field if non-nil, zero value otherwise.

### GetRestDeliveryPointNameOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetRestDeliveryPointNameOk() (*string, bool)`

GetRestDeliveryPointNameOk returns a tuple with the RestDeliveryPointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestDeliveryPointName

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetRestDeliveryPointName(v string)`

SetRestDeliveryPointName sets RestDeliveryPointName field to given value.

### HasRestDeliveryPointName

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasRestDeliveryPointName() bool`

HasRestDeliveryPointName returns a boolean if a field has been set.

### GetRetryDelay

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetRetryDelay() int32`

GetRetryDelay returns the RetryDelay field if non-nil, zero value otherwise.

### GetRetryDelayOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetRetryDelayOk() (*int32, bool)`

GetRetryDelayOk returns a tuple with the RetryDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryDelay

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetRetryDelay(v int32)`

SetRetryDelay sets RetryDelay field to given value.

### HasRetryDelay

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasRetryDelay() bool`

HasRetryDelay returns a boolean if a field has been set.

### GetTlsCipherSuiteList

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetTlsCipherSuiteList() string`

GetTlsCipherSuiteList returns the TlsCipherSuiteList field if non-nil, zero value otherwise.

### GetTlsCipherSuiteListOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetTlsCipherSuiteListOk() (*string, bool)`

GetTlsCipherSuiteListOk returns a tuple with the TlsCipherSuiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCipherSuiteList

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetTlsCipherSuiteList(v string)`

SetTlsCipherSuiteList sets TlsCipherSuiteList field to given value.

### HasTlsCipherSuiteList

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasTlsCipherSuiteList() bool`

HasTlsCipherSuiteList returns a boolean if a field has been set.

### GetTlsEnabled

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetTlsEnabled() bool`

GetTlsEnabled returns the TlsEnabled field if non-nil, zero value otherwise.

### GetTlsEnabledOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetTlsEnabledOk() (*bool, bool)`

GetTlsEnabledOk returns a tuple with the TlsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsEnabled

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetTlsEnabled(v bool)`

SetTlsEnabled sets TlsEnabled field to given value.

### HasTlsEnabled

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasTlsEnabled() bool`

HasTlsEnabled returns a boolean if a field has been set.

### GetUp

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetUp() bool`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetUpOk() (*bool, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetUp(v bool)`

SetUp sets Up field to given value.

### HasUp

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasUp() bool`

HasUp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


