# MsgVpnRestDeliveryPointRestConsumer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationClientCertContent** | Pointer to **string** | The PEM formatted content for the client certificate that the REST Consumer will present to the REST host. It must consist of a private key and between one and three certificates comprising the certificate trust chain. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. Changing this attribute requires an HTTPS connection. The default value is &#x60;\&quot;\&quot;&#x60;. Available since 2.9. | [optional] 
**AuthenticationClientCertPassword** | Pointer to **string** | The password for the client certificate. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. Changing this attribute requires an HTTPS connection. The default value is &#x60;\&quot;\&quot;&#x60;. Available since 2.9. | [optional] 
**AuthenticationHttpBasicPassword** | Pointer to **string** | The password for the username. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] 
**AuthenticationHttpBasicUsername** | Pointer to **string** | The username that the REST Consumer will use to login to the REST host. Normally a username is only configured when basic authentication is selected for the REST Consumer. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] 
**AuthenticationHttpHeaderName** | Pointer to **string** | The authentication header name. The default value is &#x60;\&quot;\&quot;&#x60;. Available since 2.15. | [optional] 
**AuthenticationHttpHeaderValue** | Pointer to **string** | The authentication header value. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. The default value is &#x60;\&quot;\&quot;&#x60;. Available since 2.15. | [optional] 
**AuthenticationOauthClientId** | Pointer to **string** | The OAuth client ID. The default value is &#x60;\&quot;\&quot;&#x60;. Available since 2.19. | [optional] 
**AuthenticationOauthClientScope** | Pointer to **string** | The OAuth scope. The default value is &#x60;\&quot;\&quot;&#x60;. Available since 2.19. | [optional] 
**AuthenticationOauthClientSecret** | Pointer to **string** | The OAuth client secret. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. The default value is &#x60;\&quot;\&quot;&#x60;. Available since 2.19. | [optional] 
**AuthenticationOauthClientTokenEndpoint** | Pointer to **string** | The OAuth token endpoint URL that the REST Consumer will use to request a token for login to the REST host. Must begin with \&quot;https\&quot;. The default value is &#x60;\&quot;\&quot;&#x60;. Available since 2.19. | [optional] 
**AuthenticationOauthJwtSecretKey** | Pointer to **string** | The OAuth secret key used to sign the token request JWT. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. The default value is &#x60;\&quot;\&quot;&#x60;. Available since 2.21. | [optional] 
**AuthenticationOauthJwtTokenEndpoint** | Pointer to **string** | The OAuth token endpoint URL that the REST Consumer will use to request a token for login to the REST host. The default value is &#x60;\&quot;\&quot;&#x60;. Available since 2.21. | [optional] 
**AuthenticationScheme** | Pointer to **string** | The authentication scheme used by the REST Consumer to login to the REST host. The default value is &#x60;\&quot;none\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;none\&quot; - Login with no authentication. This may be useful for anonymous connections or when a REST Consumer does not require authentication. \&quot;http-basic\&quot; - Login with a username and optional password according to HTTP Basic authentication as per RFC2616. \&quot;client-certificate\&quot; - Login with a client TLS certificate as per RFC5246. Client certificate authentication is only available on TLS connections. \&quot;http-header\&quot; - Login with a specified HTTP header. \&quot;oauth-client\&quot; - Login with OAuth 2.0 client credentials. \&quot;oauth-jwt\&quot; - Login with OAuth (RFC 7523 JWT Profile). \&quot;transparent\&quot; - Login using the Authorization header from the message properties, if present. Transparent authentication passes along existing Authorization header metadata instead of discarding it. Note that if the message is coming from a REST producer, the REST service must be configured to forward the Authorization header. &lt;/pre&gt;  | [optional] 
**Enabled** | Pointer to **bool** | Enable or disable the REST Consumer. When disabled, no connections are initiated or messages delivered to this particular REST Consumer. The default value is &#x60;false&#x60;. | [optional] 
**HttpMethod** | Pointer to **string** | The HTTP method to use (POST or PUT). This is used only when operating in the REST service \&quot;messaging\&quot; mode and is ignored in \&quot;gateway\&quot; mode. The default value is &#x60;\&quot;post\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;post\&quot; - Use the POST HTTP method. \&quot;put\&quot; - Use the PUT HTTP method. &lt;/pre&gt;  Available since 2.17. | [optional] 
**LocalInterface** | Pointer to **string** | The interface that will be used for all outgoing connections associated with the REST Consumer. When unspecified, an interface is automatically chosen. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] 
**MaxPostWaitTime** | Pointer to **int32** | The maximum amount of time (in seconds) to wait for an HTTP POST response from the REST Consumer. Once this time is exceeded, the TCP connection is reset. The default value is &#x60;30&#x60;. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**OutgoingConnectionCount** | Pointer to **int32** | The number of concurrent TCP connections open to the REST Consumer. The default value is &#x60;3&#x60;. | [optional] 
**RemoteHost** | Pointer to **string** | The IP address or DNS name to which the broker is to connect to deliver messages for the REST Consumer. A host value must be configured for the REST Consumer to be operationally up. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] 
**RemotePort** | Pointer to **int64** | The port associated with the host of the REST Consumer. The default value is &#x60;8080&#x60;. | [optional] 
**RestConsumerName** | Pointer to **string** | The name of the REST Consumer. | [optional] 
**RestDeliveryPointName** | Pointer to **string** | The name of the REST Delivery Point. | [optional] 
**RetryDelay** | Pointer to **int32** | The number of seconds that must pass before retrying the remote REST Consumer connection. The default value is &#x60;3&#x60;. | [optional] 
**TlsCipherSuiteList** | Pointer to **string** | The colon-separated list of cipher suites the REST Consumer uses in its encrypted connection. The value &#x60;\&quot;default\&quot;&#x60; implies all supported suites ordered from most secure to least secure. The list of default cipher suites is available in the &#x60;tlsCipherSuiteMsgBackboneDefaultList&#x60; attribute of the Broker object in the Monitoring API. The REST Consumer should choose the first suite from this list that it supports. The default value is &#x60;\&quot;default\&quot;&#x60;. | [optional] 
**TlsEnabled** | Pointer to **bool** | Enable or disable encryption (TLS) for the REST Consumer. The default value is &#x60;false&#x60;. | [optional] 

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

### GetAuthenticationClientCertContent

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationClientCertContent() string`

GetAuthenticationClientCertContent returns the AuthenticationClientCertContent field if non-nil, zero value otherwise.

### GetAuthenticationClientCertContentOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationClientCertContentOk() (*string, bool)`

GetAuthenticationClientCertContentOk returns a tuple with the AuthenticationClientCertContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationClientCertContent

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationClientCertContent(v string)`

SetAuthenticationClientCertContent sets AuthenticationClientCertContent field to given value.

### HasAuthenticationClientCertContent

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationClientCertContent() bool`

HasAuthenticationClientCertContent returns a boolean if a field has been set.

### GetAuthenticationClientCertPassword

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationClientCertPassword() string`

GetAuthenticationClientCertPassword returns the AuthenticationClientCertPassword field if non-nil, zero value otherwise.

### GetAuthenticationClientCertPasswordOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationClientCertPasswordOk() (*string, bool)`

GetAuthenticationClientCertPasswordOk returns a tuple with the AuthenticationClientCertPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationClientCertPassword

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationClientCertPassword(v string)`

SetAuthenticationClientCertPassword sets AuthenticationClientCertPassword field to given value.

### HasAuthenticationClientCertPassword

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationClientCertPassword() bool`

HasAuthenticationClientCertPassword returns a boolean if a field has been set.

### GetAuthenticationHttpBasicPassword

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationHttpBasicPassword() string`

GetAuthenticationHttpBasicPassword returns the AuthenticationHttpBasicPassword field if non-nil, zero value otherwise.

### GetAuthenticationHttpBasicPasswordOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationHttpBasicPasswordOk() (*string, bool)`

GetAuthenticationHttpBasicPasswordOk returns a tuple with the AuthenticationHttpBasicPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationHttpBasicPassword

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationHttpBasicPassword(v string)`

SetAuthenticationHttpBasicPassword sets AuthenticationHttpBasicPassword field to given value.

### HasAuthenticationHttpBasicPassword

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationHttpBasicPassword() bool`

HasAuthenticationHttpBasicPassword returns a boolean if a field has been set.

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

### GetAuthenticationHttpHeaderValue

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationHttpHeaderValue() string`

GetAuthenticationHttpHeaderValue returns the AuthenticationHttpHeaderValue field if non-nil, zero value otherwise.

### GetAuthenticationHttpHeaderValueOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationHttpHeaderValueOk() (*string, bool)`

GetAuthenticationHttpHeaderValueOk returns a tuple with the AuthenticationHttpHeaderValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationHttpHeaderValue

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationHttpHeaderValue(v string)`

SetAuthenticationHttpHeaderValue sets AuthenticationHttpHeaderValue field to given value.

### HasAuthenticationHttpHeaderValue

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationHttpHeaderValue() bool`

HasAuthenticationHttpHeaderValue returns a boolean if a field has been set.

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

### GetAuthenticationOauthClientSecret

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthClientSecret() string`

GetAuthenticationOauthClientSecret returns the AuthenticationOauthClientSecret field if non-nil, zero value otherwise.

### GetAuthenticationOauthClientSecretOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthClientSecretOk() (*string, bool)`

GetAuthenticationOauthClientSecretOk returns a tuple with the AuthenticationOauthClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationOauthClientSecret

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationOauthClientSecret(v string)`

SetAuthenticationOauthClientSecret sets AuthenticationOauthClientSecret field to given value.

### HasAuthenticationOauthClientSecret

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationOauthClientSecret() bool`

HasAuthenticationOauthClientSecret returns a boolean if a field has been set.

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

### GetAuthenticationOauthJwtSecretKey

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthJwtSecretKey() string`

GetAuthenticationOauthJwtSecretKey returns the AuthenticationOauthJwtSecretKey field if non-nil, zero value otherwise.

### GetAuthenticationOauthJwtSecretKeyOk

`func (o *MsgVpnRestDeliveryPointRestConsumer) GetAuthenticationOauthJwtSecretKeyOk() (*string, bool)`

GetAuthenticationOauthJwtSecretKeyOk returns a tuple with the AuthenticationOauthJwtSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationOauthJwtSecretKey

`func (o *MsgVpnRestDeliveryPointRestConsumer) SetAuthenticationOauthJwtSecretKey(v string)`

SetAuthenticationOauthJwtSecretKey sets AuthenticationOauthJwtSecretKey field to given value.

### HasAuthenticationOauthJwtSecretKey

`func (o *MsgVpnRestDeliveryPointRestConsumer) HasAuthenticationOauthJwtSecretKey() bool`

HasAuthenticationOauthJwtSecretKey returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


