# MsgVpnAuthenticationOauthProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudienceClaimName** | Pointer to **string** | The audience claim name, indicating which part of the object to use for determining the audience. | [optional] 
**AudienceClaimSource** | Pointer to **string** | The audience claim source, indicating where to search for the audience value. The allowed values and their meaning are:  &lt;pre&gt; \&quot;access-token\&quot; - The OAuth v2 access_token. \&quot;id-token\&quot; - The OpenID Connect id_token. \&quot;introspection\&quot; - The result of introspecting the OAuth v2 access_token. &lt;/pre&gt;  | [optional] 
**AudienceClaimValue** | Pointer to **string** | The required audience value for a token to be considered valid. | [optional] 
**AudienceValidationEnabled** | Pointer to **bool** | Indicates whether audience validation is enabled. | [optional] 
**AuthenticationSuccessCount** | Pointer to **int64** | The number of OAuth Provider client authentications that succeeded. | [optional] 
**AuthorizationGroupClaimName** | Pointer to **string** | The authorization group claim name, indicating which part of the object to use for determining the authorization group. | [optional] 
**AuthorizationGroupClaimSource** | Pointer to **string** | The authorization group claim source, indicating where to search for the authorization group name. The allowed values and their meaning are:  &lt;pre&gt; \&quot;access-token\&quot; - The OAuth v2 access_token. \&quot;id-token\&quot; - The OpenID Connect id_token. \&quot;introspection\&quot; - The result of introspecting the OAuth v2 access_token. &lt;/pre&gt;  | [optional] 
**AuthorizationGroupEnabled** | Pointer to **bool** | Indicates whether OAuth based authorization is enabled and the configured authorization type for OAuth clients is overridden. | [optional] 
**DisconnectOnTokenExpirationEnabled** | Pointer to **bool** | Indicates whether clients are disconnected when their tokens expire. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether OAuth Provider client authentication is enabled. | [optional] 
**JwksLastRefreshFailureReason** | Pointer to **string** | The reason for the last JWKS public key refresh failure. | [optional] 
**JwksLastRefreshFailureTime** | Pointer to **int32** | The timestamp of the last JWKS public key refresh failure. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] 
**JwksLastRefreshTime** | Pointer to **int32** | The timestamp of the last JWKS public key refresh success. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] 
**JwksNextScheduledRefreshTime** | Pointer to **int32** | The timestamp of the next scheduled JWKS public key refresh. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] 
**JwksRefreshFailureCount** | Pointer to **int64** | The number of JWKS public key refresh failures. | [optional] 
**JwksRefreshInterval** | Pointer to **int32** | The number of seconds between forced JWKS public key refreshing. | [optional] 
**JwksUri** | Pointer to **string** | The URI where the OAuth provider publishes its JWKS public keys. | [optional] 
**LoginFailureIncorrectAudienceValueCount** | Pointer to **int64** | The number of login failures due to an incorrect audience value. | [optional] 
**LoginFailureInvalidAudienceValueCount** | Pointer to **int64** | The number of login failures due to an invalid audience value. | [optional] 
**LoginFailureInvalidAuthorizationGroupValueCount** | Pointer to **int64** | The number of login failures due to an invalid authorization group value (zero-length or non-string). | [optional] 
**LoginFailureInvalidJwtSignatureCount** | Pointer to **int64** | The number of login failures due to an invalid JWT signature. | [optional] 
**LoginFailureInvalidUsernameValueCount** | Pointer to **int64** | The number of login failures due to an invalid username value. | [optional] 
**LoginFailureMismatchedUsernameCount** | Pointer to **int64** | The number of login failures due to a mismatched username. | [optional] 
**LoginFailureMissingAudienceCount** | Pointer to **int64** | The number of login failures due to a missing audience claim. | [optional] 
**LoginFailureMissingJwkCount** | Pointer to **int64** | The number of login failures due to a missing JSON Web Key (JWK). | [optional] 
**LoginFailureMissingOrInvalidTokenCount** | Pointer to **int64** | The number of login failures due to a missing or invalid token. | [optional] 
**LoginFailureMissingUsernameCount** | Pointer to **int64** | The number of login failures due to a missing username. | [optional] 
**LoginFailureTokenExpiredCount** | Pointer to **int64** | The number of login failures due to a token being expired. | [optional] 
**LoginFailureTokenIntrospectionErroredCount** | Pointer to **int64** | The number of login failures due to a token introspection error response. | [optional] 
**LoginFailureTokenIntrospectionFailureCount** | Pointer to **int64** | The number of login failures due to a failure to complete the token introspection. | [optional] 
**LoginFailureTokenIntrospectionHttpsErrorCount** | Pointer to **int64** | The number of login failures due to a token introspection HTTPS error. | [optional] 
**LoginFailureTokenIntrospectionInvalidCount** | Pointer to **int64** | The number of login failures due to a token introspection response being invalid. | [optional] 
**LoginFailureTokenIntrospectionTimeoutCount** | Pointer to **int64** | The number of login failures due to a token introspection timeout. | [optional] 
**LoginFailureTokenNotValidYetCount** | Pointer to **int64** | The number of login failures due to a token not being valid yet. | [optional] 
**LoginFailureUnsupportedAlgCount** | Pointer to **int64** | The number of login failures due to an unsupported algorithm. | [optional] 
**MissingAuthorizationGroupCount** | Pointer to **int64** | The number of clients that did not provide an authorization group claim value when expected. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**OauthProviderName** | Pointer to **string** | The name of the OAuth Provider. | [optional] 
**TokenIgnoreTimeLimitsEnabled** | Pointer to **bool** | Indicates whether to ignore time limits and accept tokens that are not yet valid or are no longer valid. | [optional] 
**TokenIntrospectionAverageTime** | Pointer to **int32** | The one minute average of the time required to complete a token introspection, in milliseconds (ms). | [optional] 
**TokenIntrospectionLastFailureReason** | Pointer to **string** | The reason for the last token introspection failure. | [optional] 
**TokenIntrospectionLastFailureTime** | Pointer to **int32** | The timestamp of the last token introspection failure. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] 
**TokenIntrospectionParameterName** | Pointer to **string** | The parameter name used to identify the token during access token introspection. A standards compliant OAuth introspection server expects \&quot;token\&quot;. | [optional] 
**TokenIntrospectionSuccessCount** | Pointer to **int64** | The number of token introspection successes. | [optional] 
**TokenIntrospectionTimeout** | Pointer to **int32** | The maximum time in seconds a token introspection is allowed to take. | [optional] 
**TokenIntrospectionUri** | Pointer to **string** | The token introspection URI of the OAuth authentication server. | [optional] 
**TokenIntrospectionUsername** | Pointer to **string** | The username to use when logging into the token introspection URI. | [optional] 
**UsernameClaimName** | Pointer to **string** | The username claim name, indicating which part of the object to use for determining the username. | [optional] 
**UsernameClaimSource** | Pointer to **string** | The username claim source, indicating where to search for the username value. The allowed values and their meaning are:  &lt;pre&gt; \&quot;access-token\&quot; - The OAuth v2 access_token. \&quot;id-token\&quot; - The OpenID Connect id_token. \&quot;introspection\&quot; - The result of introspecting the OAuth v2 access_token. &lt;/pre&gt;  | [optional] 
**UsernameValidateEnabled** | Pointer to **bool** | Indicates whether the API provided username will be validated against the username calculated from the token(s). | [optional] 

## Methods

### NewMsgVpnAuthenticationOauthProvider

`func NewMsgVpnAuthenticationOauthProvider() *MsgVpnAuthenticationOauthProvider`

NewMsgVpnAuthenticationOauthProvider instantiates a new MsgVpnAuthenticationOauthProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnAuthenticationOauthProviderWithDefaults

`func NewMsgVpnAuthenticationOauthProviderWithDefaults() *MsgVpnAuthenticationOauthProvider`

NewMsgVpnAuthenticationOauthProviderWithDefaults instantiates a new MsgVpnAuthenticationOauthProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudienceClaimName

`func (o *MsgVpnAuthenticationOauthProvider) GetAudienceClaimName() string`

GetAudienceClaimName returns the AudienceClaimName field if non-nil, zero value otherwise.

### GetAudienceClaimNameOk

`func (o *MsgVpnAuthenticationOauthProvider) GetAudienceClaimNameOk() (*string, bool)`

GetAudienceClaimNameOk returns a tuple with the AudienceClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceClaimName

`func (o *MsgVpnAuthenticationOauthProvider) SetAudienceClaimName(v string)`

SetAudienceClaimName sets AudienceClaimName field to given value.

### HasAudienceClaimName

`func (o *MsgVpnAuthenticationOauthProvider) HasAudienceClaimName() bool`

HasAudienceClaimName returns a boolean if a field has been set.

### GetAudienceClaimSource

`func (o *MsgVpnAuthenticationOauthProvider) GetAudienceClaimSource() string`

GetAudienceClaimSource returns the AudienceClaimSource field if non-nil, zero value otherwise.

### GetAudienceClaimSourceOk

`func (o *MsgVpnAuthenticationOauthProvider) GetAudienceClaimSourceOk() (*string, bool)`

GetAudienceClaimSourceOk returns a tuple with the AudienceClaimSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceClaimSource

`func (o *MsgVpnAuthenticationOauthProvider) SetAudienceClaimSource(v string)`

SetAudienceClaimSource sets AudienceClaimSource field to given value.

### HasAudienceClaimSource

`func (o *MsgVpnAuthenticationOauthProvider) HasAudienceClaimSource() bool`

HasAudienceClaimSource returns a boolean if a field has been set.

### GetAudienceClaimValue

`func (o *MsgVpnAuthenticationOauthProvider) GetAudienceClaimValue() string`

GetAudienceClaimValue returns the AudienceClaimValue field if non-nil, zero value otherwise.

### GetAudienceClaimValueOk

`func (o *MsgVpnAuthenticationOauthProvider) GetAudienceClaimValueOk() (*string, bool)`

GetAudienceClaimValueOk returns a tuple with the AudienceClaimValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceClaimValue

`func (o *MsgVpnAuthenticationOauthProvider) SetAudienceClaimValue(v string)`

SetAudienceClaimValue sets AudienceClaimValue field to given value.

### HasAudienceClaimValue

`func (o *MsgVpnAuthenticationOauthProvider) HasAudienceClaimValue() bool`

HasAudienceClaimValue returns a boolean if a field has been set.

### GetAudienceValidationEnabled

`func (o *MsgVpnAuthenticationOauthProvider) GetAudienceValidationEnabled() bool`

GetAudienceValidationEnabled returns the AudienceValidationEnabled field if non-nil, zero value otherwise.

### GetAudienceValidationEnabledOk

`func (o *MsgVpnAuthenticationOauthProvider) GetAudienceValidationEnabledOk() (*bool, bool)`

GetAudienceValidationEnabledOk returns a tuple with the AudienceValidationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceValidationEnabled

`func (o *MsgVpnAuthenticationOauthProvider) SetAudienceValidationEnabled(v bool)`

SetAudienceValidationEnabled sets AudienceValidationEnabled field to given value.

### HasAudienceValidationEnabled

`func (o *MsgVpnAuthenticationOauthProvider) HasAudienceValidationEnabled() bool`

HasAudienceValidationEnabled returns a boolean if a field has been set.

### GetAuthenticationSuccessCount

`func (o *MsgVpnAuthenticationOauthProvider) GetAuthenticationSuccessCount() int64`

GetAuthenticationSuccessCount returns the AuthenticationSuccessCount field if non-nil, zero value otherwise.

### GetAuthenticationSuccessCountOk

`func (o *MsgVpnAuthenticationOauthProvider) GetAuthenticationSuccessCountOk() (*int64, bool)`

GetAuthenticationSuccessCountOk returns a tuple with the AuthenticationSuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationSuccessCount

`func (o *MsgVpnAuthenticationOauthProvider) SetAuthenticationSuccessCount(v int64)`

SetAuthenticationSuccessCount sets AuthenticationSuccessCount field to given value.

### HasAuthenticationSuccessCount

`func (o *MsgVpnAuthenticationOauthProvider) HasAuthenticationSuccessCount() bool`

HasAuthenticationSuccessCount returns a boolean if a field has been set.

### GetAuthorizationGroupClaimName

`func (o *MsgVpnAuthenticationOauthProvider) GetAuthorizationGroupClaimName() string`

GetAuthorizationGroupClaimName returns the AuthorizationGroupClaimName field if non-nil, zero value otherwise.

### GetAuthorizationGroupClaimNameOk

`func (o *MsgVpnAuthenticationOauthProvider) GetAuthorizationGroupClaimNameOk() (*string, bool)`

GetAuthorizationGroupClaimNameOk returns a tuple with the AuthorizationGroupClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationGroupClaimName

`func (o *MsgVpnAuthenticationOauthProvider) SetAuthorizationGroupClaimName(v string)`

SetAuthorizationGroupClaimName sets AuthorizationGroupClaimName field to given value.

### HasAuthorizationGroupClaimName

`func (o *MsgVpnAuthenticationOauthProvider) HasAuthorizationGroupClaimName() bool`

HasAuthorizationGroupClaimName returns a boolean if a field has been set.

### GetAuthorizationGroupClaimSource

`func (o *MsgVpnAuthenticationOauthProvider) GetAuthorizationGroupClaimSource() string`

GetAuthorizationGroupClaimSource returns the AuthorizationGroupClaimSource field if non-nil, zero value otherwise.

### GetAuthorizationGroupClaimSourceOk

`func (o *MsgVpnAuthenticationOauthProvider) GetAuthorizationGroupClaimSourceOk() (*string, bool)`

GetAuthorizationGroupClaimSourceOk returns a tuple with the AuthorizationGroupClaimSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationGroupClaimSource

`func (o *MsgVpnAuthenticationOauthProvider) SetAuthorizationGroupClaimSource(v string)`

SetAuthorizationGroupClaimSource sets AuthorizationGroupClaimSource field to given value.

### HasAuthorizationGroupClaimSource

`func (o *MsgVpnAuthenticationOauthProvider) HasAuthorizationGroupClaimSource() bool`

HasAuthorizationGroupClaimSource returns a boolean if a field has been set.

### GetAuthorizationGroupEnabled

`func (o *MsgVpnAuthenticationOauthProvider) GetAuthorizationGroupEnabled() bool`

GetAuthorizationGroupEnabled returns the AuthorizationGroupEnabled field if non-nil, zero value otherwise.

### GetAuthorizationGroupEnabledOk

`func (o *MsgVpnAuthenticationOauthProvider) GetAuthorizationGroupEnabledOk() (*bool, bool)`

GetAuthorizationGroupEnabledOk returns a tuple with the AuthorizationGroupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationGroupEnabled

`func (o *MsgVpnAuthenticationOauthProvider) SetAuthorizationGroupEnabled(v bool)`

SetAuthorizationGroupEnabled sets AuthorizationGroupEnabled field to given value.

### HasAuthorizationGroupEnabled

`func (o *MsgVpnAuthenticationOauthProvider) HasAuthorizationGroupEnabled() bool`

HasAuthorizationGroupEnabled returns a boolean if a field has been set.

### GetDisconnectOnTokenExpirationEnabled

`func (o *MsgVpnAuthenticationOauthProvider) GetDisconnectOnTokenExpirationEnabled() bool`

GetDisconnectOnTokenExpirationEnabled returns the DisconnectOnTokenExpirationEnabled field if non-nil, zero value otherwise.

### GetDisconnectOnTokenExpirationEnabledOk

`func (o *MsgVpnAuthenticationOauthProvider) GetDisconnectOnTokenExpirationEnabledOk() (*bool, bool)`

GetDisconnectOnTokenExpirationEnabledOk returns a tuple with the DisconnectOnTokenExpirationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectOnTokenExpirationEnabled

`func (o *MsgVpnAuthenticationOauthProvider) SetDisconnectOnTokenExpirationEnabled(v bool)`

SetDisconnectOnTokenExpirationEnabled sets DisconnectOnTokenExpirationEnabled field to given value.

### HasDisconnectOnTokenExpirationEnabled

`func (o *MsgVpnAuthenticationOauthProvider) HasDisconnectOnTokenExpirationEnabled() bool`

HasDisconnectOnTokenExpirationEnabled returns a boolean if a field has been set.

### GetEnabled

`func (o *MsgVpnAuthenticationOauthProvider) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MsgVpnAuthenticationOauthProvider) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MsgVpnAuthenticationOauthProvider) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MsgVpnAuthenticationOauthProvider) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetJwksLastRefreshFailureReason

`func (o *MsgVpnAuthenticationOauthProvider) GetJwksLastRefreshFailureReason() string`

GetJwksLastRefreshFailureReason returns the JwksLastRefreshFailureReason field if non-nil, zero value otherwise.

### GetJwksLastRefreshFailureReasonOk

`func (o *MsgVpnAuthenticationOauthProvider) GetJwksLastRefreshFailureReasonOk() (*string, bool)`

GetJwksLastRefreshFailureReasonOk returns a tuple with the JwksLastRefreshFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksLastRefreshFailureReason

`func (o *MsgVpnAuthenticationOauthProvider) SetJwksLastRefreshFailureReason(v string)`

SetJwksLastRefreshFailureReason sets JwksLastRefreshFailureReason field to given value.

### HasJwksLastRefreshFailureReason

`func (o *MsgVpnAuthenticationOauthProvider) HasJwksLastRefreshFailureReason() bool`

HasJwksLastRefreshFailureReason returns a boolean if a field has been set.

### GetJwksLastRefreshFailureTime

`func (o *MsgVpnAuthenticationOauthProvider) GetJwksLastRefreshFailureTime() int32`

GetJwksLastRefreshFailureTime returns the JwksLastRefreshFailureTime field if non-nil, zero value otherwise.

### GetJwksLastRefreshFailureTimeOk

`func (o *MsgVpnAuthenticationOauthProvider) GetJwksLastRefreshFailureTimeOk() (*int32, bool)`

GetJwksLastRefreshFailureTimeOk returns a tuple with the JwksLastRefreshFailureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksLastRefreshFailureTime

`func (o *MsgVpnAuthenticationOauthProvider) SetJwksLastRefreshFailureTime(v int32)`

SetJwksLastRefreshFailureTime sets JwksLastRefreshFailureTime field to given value.

### HasJwksLastRefreshFailureTime

`func (o *MsgVpnAuthenticationOauthProvider) HasJwksLastRefreshFailureTime() bool`

HasJwksLastRefreshFailureTime returns a boolean if a field has been set.

### GetJwksLastRefreshTime

`func (o *MsgVpnAuthenticationOauthProvider) GetJwksLastRefreshTime() int32`

GetJwksLastRefreshTime returns the JwksLastRefreshTime field if non-nil, zero value otherwise.

### GetJwksLastRefreshTimeOk

`func (o *MsgVpnAuthenticationOauthProvider) GetJwksLastRefreshTimeOk() (*int32, bool)`

GetJwksLastRefreshTimeOk returns a tuple with the JwksLastRefreshTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksLastRefreshTime

`func (o *MsgVpnAuthenticationOauthProvider) SetJwksLastRefreshTime(v int32)`

SetJwksLastRefreshTime sets JwksLastRefreshTime field to given value.

### HasJwksLastRefreshTime

`func (o *MsgVpnAuthenticationOauthProvider) HasJwksLastRefreshTime() bool`

HasJwksLastRefreshTime returns a boolean if a field has been set.

### GetJwksNextScheduledRefreshTime

`func (o *MsgVpnAuthenticationOauthProvider) GetJwksNextScheduledRefreshTime() int32`

GetJwksNextScheduledRefreshTime returns the JwksNextScheduledRefreshTime field if non-nil, zero value otherwise.

### GetJwksNextScheduledRefreshTimeOk

`func (o *MsgVpnAuthenticationOauthProvider) GetJwksNextScheduledRefreshTimeOk() (*int32, bool)`

GetJwksNextScheduledRefreshTimeOk returns a tuple with the JwksNextScheduledRefreshTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksNextScheduledRefreshTime

`func (o *MsgVpnAuthenticationOauthProvider) SetJwksNextScheduledRefreshTime(v int32)`

SetJwksNextScheduledRefreshTime sets JwksNextScheduledRefreshTime field to given value.

### HasJwksNextScheduledRefreshTime

`func (o *MsgVpnAuthenticationOauthProvider) HasJwksNextScheduledRefreshTime() bool`

HasJwksNextScheduledRefreshTime returns a boolean if a field has been set.

### GetJwksRefreshFailureCount

`func (o *MsgVpnAuthenticationOauthProvider) GetJwksRefreshFailureCount() int64`

GetJwksRefreshFailureCount returns the JwksRefreshFailureCount field if non-nil, zero value otherwise.

### GetJwksRefreshFailureCountOk

`func (o *MsgVpnAuthenticationOauthProvider) GetJwksRefreshFailureCountOk() (*int64, bool)`

GetJwksRefreshFailureCountOk returns a tuple with the JwksRefreshFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksRefreshFailureCount

`func (o *MsgVpnAuthenticationOauthProvider) SetJwksRefreshFailureCount(v int64)`

SetJwksRefreshFailureCount sets JwksRefreshFailureCount field to given value.

### HasJwksRefreshFailureCount

`func (o *MsgVpnAuthenticationOauthProvider) HasJwksRefreshFailureCount() bool`

HasJwksRefreshFailureCount returns a boolean if a field has been set.

### GetJwksRefreshInterval

`func (o *MsgVpnAuthenticationOauthProvider) GetJwksRefreshInterval() int32`

GetJwksRefreshInterval returns the JwksRefreshInterval field if non-nil, zero value otherwise.

### GetJwksRefreshIntervalOk

`func (o *MsgVpnAuthenticationOauthProvider) GetJwksRefreshIntervalOk() (*int32, bool)`

GetJwksRefreshIntervalOk returns a tuple with the JwksRefreshInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksRefreshInterval

`func (o *MsgVpnAuthenticationOauthProvider) SetJwksRefreshInterval(v int32)`

SetJwksRefreshInterval sets JwksRefreshInterval field to given value.

### HasJwksRefreshInterval

`func (o *MsgVpnAuthenticationOauthProvider) HasJwksRefreshInterval() bool`

HasJwksRefreshInterval returns a boolean if a field has been set.

### GetJwksUri

`func (o *MsgVpnAuthenticationOauthProvider) GetJwksUri() string`

GetJwksUri returns the JwksUri field if non-nil, zero value otherwise.

### GetJwksUriOk

`func (o *MsgVpnAuthenticationOauthProvider) GetJwksUriOk() (*string, bool)`

GetJwksUriOk returns a tuple with the JwksUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksUri

`func (o *MsgVpnAuthenticationOauthProvider) SetJwksUri(v string)`

SetJwksUri sets JwksUri field to given value.

### HasJwksUri

`func (o *MsgVpnAuthenticationOauthProvider) HasJwksUri() bool`

HasJwksUri returns a boolean if a field has been set.

### GetLoginFailureIncorrectAudienceValueCount

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureIncorrectAudienceValueCount() int64`

GetLoginFailureIncorrectAudienceValueCount returns the LoginFailureIncorrectAudienceValueCount field if non-nil, zero value otherwise.

### GetLoginFailureIncorrectAudienceValueCountOk

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureIncorrectAudienceValueCountOk() (*int64, bool)`

GetLoginFailureIncorrectAudienceValueCountOk returns a tuple with the LoginFailureIncorrectAudienceValueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginFailureIncorrectAudienceValueCount

`func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureIncorrectAudienceValueCount(v int64)`

SetLoginFailureIncorrectAudienceValueCount sets LoginFailureIncorrectAudienceValueCount field to given value.

### HasLoginFailureIncorrectAudienceValueCount

`func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureIncorrectAudienceValueCount() bool`

HasLoginFailureIncorrectAudienceValueCount returns a boolean if a field has been set.

### GetLoginFailureInvalidAudienceValueCount

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureInvalidAudienceValueCount() int64`

GetLoginFailureInvalidAudienceValueCount returns the LoginFailureInvalidAudienceValueCount field if non-nil, zero value otherwise.

### GetLoginFailureInvalidAudienceValueCountOk

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureInvalidAudienceValueCountOk() (*int64, bool)`

GetLoginFailureInvalidAudienceValueCountOk returns a tuple with the LoginFailureInvalidAudienceValueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginFailureInvalidAudienceValueCount

`func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureInvalidAudienceValueCount(v int64)`

SetLoginFailureInvalidAudienceValueCount sets LoginFailureInvalidAudienceValueCount field to given value.

### HasLoginFailureInvalidAudienceValueCount

`func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureInvalidAudienceValueCount() bool`

HasLoginFailureInvalidAudienceValueCount returns a boolean if a field has been set.

### GetLoginFailureInvalidAuthorizationGroupValueCount

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureInvalidAuthorizationGroupValueCount() int64`

GetLoginFailureInvalidAuthorizationGroupValueCount returns the LoginFailureInvalidAuthorizationGroupValueCount field if non-nil, zero value otherwise.

### GetLoginFailureInvalidAuthorizationGroupValueCountOk

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureInvalidAuthorizationGroupValueCountOk() (*int64, bool)`

GetLoginFailureInvalidAuthorizationGroupValueCountOk returns a tuple with the LoginFailureInvalidAuthorizationGroupValueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginFailureInvalidAuthorizationGroupValueCount

`func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureInvalidAuthorizationGroupValueCount(v int64)`

SetLoginFailureInvalidAuthorizationGroupValueCount sets LoginFailureInvalidAuthorizationGroupValueCount field to given value.

### HasLoginFailureInvalidAuthorizationGroupValueCount

`func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureInvalidAuthorizationGroupValueCount() bool`

HasLoginFailureInvalidAuthorizationGroupValueCount returns a boolean if a field has been set.

### GetLoginFailureInvalidJwtSignatureCount

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureInvalidJwtSignatureCount() int64`

GetLoginFailureInvalidJwtSignatureCount returns the LoginFailureInvalidJwtSignatureCount field if non-nil, zero value otherwise.

### GetLoginFailureInvalidJwtSignatureCountOk

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureInvalidJwtSignatureCountOk() (*int64, bool)`

GetLoginFailureInvalidJwtSignatureCountOk returns a tuple with the LoginFailureInvalidJwtSignatureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginFailureInvalidJwtSignatureCount

`func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureInvalidJwtSignatureCount(v int64)`

SetLoginFailureInvalidJwtSignatureCount sets LoginFailureInvalidJwtSignatureCount field to given value.

### HasLoginFailureInvalidJwtSignatureCount

`func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureInvalidJwtSignatureCount() bool`

HasLoginFailureInvalidJwtSignatureCount returns a boolean if a field has been set.

### GetLoginFailureInvalidUsernameValueCount

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureInvalidUsernameValueCount() int64`

GetLoginFailureInvalidUsernameValueCount returns the LoginFailureInvalidUsernameValueCount field if non-nil, zero value otherwise.

### GetLoginFailureInvalidUsernameValueCountOk

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureInvalidUsernameValueCountOk() (*int64, bool)`

GetLoginFailureInvalidUsernameValueCountOk returns a tuple with the LoginFailureInvalidUsernameValueCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginFailureInvalidUsernameValueCount

`func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureInvalidUsernameValueCount(v int64)`

SetLoginFailureInvalidUsernameValueCount sets LoginFailureInvalidUsernameValueCount field to given value.

### HasLoginFailureInvalidUsernameValueCount

`func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureInvalidUsernameValueCount() bool`

HasLoginFailureInvalidUsernameValueCount returns a boolean if a field has been set.

### GetLoginFailureMismatchedUsernameCount

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureMismatchedUsernameCount() int64`

GetLoginFailureMismatchedUsernameCount returns the LoginFailureMismatchedUsernameCount field if non-nil, zero value otherwise.

### GetLoginFailureMismatchedUsernameCountOk

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureMismatchedUsernameCountOk() (*int64, bool)`

GetLoginFailureMismatchedUsernameCountOk returns a tuple with the LoginFailureMismatchedUsernameCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginFailureMismatchedUsernameCount

`func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureMismatchedUsernameCount(v int64)`

SetLoginFailureMismatchedUsernameCount sets LoginFailureMismatchedUsernameCount field to given value.

### HasLoginFailureMismatchedUsernameCount

`func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureMismatchedUsernameCount() bool`

HasLoginFailureMismatchedUsernameCount returns a boolean if a field has been set.

### GetLoginFailureMissingAudienceCount

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureMissingAudienceCount() int64`

GetLoginFailureMissingAudienceCount returns the LoginFailureMissingAudienceCount field if non-nil, zero value otherwise.

### GetLoginFailureMissingAudienceCountOk

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureMissingAudienceCountOk() (*int64, bool)`

GetLoginFailureMissingAudienceCountOk returns a tuple with the LoginFailureMissingAudienceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginFailureMissingAudienceCount

`func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureMissingAudienceCount(v int64)`

SetLoginFailureMissingAudienceCount sets LoginFailureMissingAudienceCount field to given value.

### HasLoginFailureMissingAudienceCount

`func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureMissingAudienceCount() bool`

HasLoginFailureMissingAudienceCount returns a boolean if a field has been set.

### GetLoginFailureMissingJwkCount

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureMissingJwkCount() int64`

GetLoginFailureMissingJwkCount returns the LoginFailureMissingJwkCount field if non-nil, zero value otherwise.

### GetLoginFailureMissingJwkCountOk

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureMissingJwkCountOk() (*int64, bool)`

GetLoginFailureMissingJwkCountOk returns a tuple with the LoginFailureMissingJwkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginFailureMissingJwkCount

`func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureMissingJwkCount(v int64)`

SetLoginFailureMissingJwkCount sets LoginFailureMissingJwkCount field to given value.

### HasLoginFailureMissingJwkCount

`func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureMissingJwkCount() bool`

HasLoginFailureMissingJwkCount returns a boolean if a field has been set.

### GetLoginFailureMissingOrInvalidTokenCount

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureMissingOrInvalidTokenCount() int64`

GetLoginFailureMissingOrInvalidTokenCount returns the LoginFailureMissingOrInvalidTokenCount field if non-nil, zero value otherwise.

### GetLoginFailureMissingOrInvalidTokenCountOk

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureMissingOrInvalidTokenCountOk() (*int64, bool)`

GetLoginFailureMissingOrInvalidTokenCountOk returns a tuple with the LoginFailureMissingOrInvalidTokenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginFailureMissingOrInvalidTokenCount

`func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureMissingOrInvalidTokenCount(v int64)`

SetLoginFailureMissingOrInvalidTokenCount sets LoginFailureMissingOrInvalidTokenCount field to given value.

### HasLoginFailureMissingOrInvalidTokenCount

`func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureMissingOrInvalidTokenCount() bool`

HasLoginFailureMissingOrInvalidTokenCount returns a boolean if a field has been set.

### GetLoginFailureMissingUsernameCount

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureMissingUsernameCount() int64`

GetLoginFailureMissingUsernameCount returns the LoginFailureMissingUsernameCount field if non-nil, zero value otherwise.

### GetLoginFailureMissingUsernameCountOk

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureMissingUsernameCountOk() (*int64, bool)`

GetLoginFailureMissingUsernameCountOk returns a tuple with the LoginFailureMissingUsernameCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginFailureMissingUsernameCount

`func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureMissingUsernameCount(v int64)`

SetLoginFailureMissingUsernameCount sets LoginFailureMissingUsernameCount field to given value.

### HasLoginFailureMissingUsernameCount

`func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureMissingUsernameCount() bool`

HasLoginFailureMissingUsernameCount returns a boolean if a field has been set.

### GetLoginFailureTokenExpiredCount

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureTokenExpiredCount() int64`

GetLoginFailureTokenExpiredCount returns the LoginFailureTokenExpiredCount field if non-nil, zero value otherwise.

### GetLoginFailureTokenExpiredCountOk

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureTokenExpiredCountOk() (*int64, bool)`

GetLoginFailureTokenExpiredCountOk returns a tuple with the LoginFailureTokenExpiredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginFailureTokenExpiredCount

`func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureTokenExpiredCount(v int64)`

SetLoginFailureTokenExpiredCount sets LoginFailureTokenExpiredCount field to given value.

### HasLoginFailureTokenExpiredCount

`func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureTokenExpiredCount() bool`

HasLoginFailureTokenExpiredCount returns a boolean if a field has been set.

### GetLoginFailureTokenIntrospectionErroredCount

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureTokenIntrospectionErroredCount() int64`

GetLoginFailureTokenIntrospectionErroredCount returns the LoginFailureTokenIntrospectionErroredCount field if non-nil, zero value otherwise.

### GetLoginFailureTokenIntrospectionErroredCountOk

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureTokenIntrospectionErroredCountOk() (*int64, bool)`

GetLoginFailureTokenIntrospectionErroredCountOk returns a tuple with the LoginFailureTokenIntrospectionErroredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginFailureTokenIntrospectionErroredCount

`func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureTokenIntrospectionErroredCount(v int64)`

SetLoginFailureTokenIntrospectionErroredCount sets LoginFailureTokenIntrospectionErroredCount field to given value.

### HasLoginFailureTokenIntrospectionErroredCount

`func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureTokenIntrospectionErroredCount() bool`

HasLoginFailureTokenIntrospectionErroredCount returns a boolean if a field has been set.

### GetLoginFailureTokenIntrospectionFailureCount

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureTokenIntrospectionFailureCount() int64`

GetLoginFailureTokenIntrospectionFailureCount returns the LoginFailureTokenIntrospectionFailureCount field if non-nil, zero value otherwise.

### GetLoginFailureTokenIntrospectionFailureCountOk

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureTokenIntrospectionFailureCountOk() (*int64, bool)`

GetLoginFailureTokenIntrospectionFailureCountOk returns a tuple with the LoginFailureTokenIntrospectionFailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginFailureTokenIntrospectionFailureCount

`func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureTokenIntrospectionFailureCount(v int64)`

SetLoginFailureTokenIntrospectionFailureCount sets LoginFailureTokenIntrospectionFailureCount field to given value.

### HasLoginFailureTokenIntrospectionFailureCount

`func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureTokenIntrospectionFailureCount() bool`

HasLoginFailureTokenIntrospectionFailureCount returns a boolean if a field has been set.

### GetLoginFailureTokenIntrospectionHttpsErrorCount

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureTokenIntrospectionHttpsErrorCount() int64`

GetLoginFailureTokenIntrospectionHttpsErrorCount returns the LoginFailureTokenIntrospectionHttpsErrorCount field if non-nil, zero value otherwise.

### GetLoginFailureTokenIntrospectionHttpsErrorCountOk

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureTokenIntrospectionHttpsErrorCountOk() (*int64, bool)`

GetLoginFailureTokenIntrospectionHttpsErrorCountOk returns a tuple with the LoginFailureTokenIntrospectionHttpsErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginFailureTokenIntrospectionHttpsErrorCount

`func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureTokenIntrospectionHttpsErrorCount(v int64)`

SetLoginFailureTokenIntrospectionHttpsErrorCount sets LoginFailureTokenIntrospectionHttpsErrorCount field to given value.

### HasLoginFailureTokenIntrospectionHttpsErrorCount

`func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureTokenIntrospectionHttpsErrorCount() bool`

HasLoginFailureTokenIntrospectionHttpsErrorCount returns a boolean if a field has been set.

### GetLoginFailureTokenIntrospectionInvalidCount

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureTokenIntrospectionInvalidCount() int64`

GetLoginFailureTokenIntrospectionInvalidCount returns the LoginFailureTokenIntrospectionInvalidCount field if non-nil, zero value otherwise.

### GetLoginFailureTokenIntrospectionInvalidCountOk

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureTokenIntrospectionInvalidCountOk() (*int64, bool)`

GetLoginFailureTokenIntrospectionInvalidCountOk returns a tuple with the LoginFailureTokenIntrospectionInvalidCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginFailureTokenIntrospectionInvalidCount

`func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureTokenIntrospectionInvalidCount(v int64)`

SetLoginFailureTokenIntrospectionInvalidCount sets LoginFailureTokenIntrospectionInvalidCount field to given value.

### HasLoginFailureTokenIntrospectionInvalidCount

`func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureTokenIntrospectionInvalidCount() bool`

HasLoginFailureTokenIntrospectionInvalidCount returns a boolean if a field has been set.

### GetLoginFailureTokenIntrospectionTimeoutCount

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureTokenIntrospectionTimeoutCount() int64`

GetLoginFailureTokenIntrospectionTimeoutCount returns the LoginFailureTokenIntrospectionTimeoutCount field if non-nil, zero value otherwise.

### GetLoginFailureTokenIntrospectionTimeoutCountOk

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureTokenIntrospectionTimeoutCountOk() (*int64, bool)`

GetLoginFailureTokenIntrospectionTimeoutCountOk returns a tuple with the LoginFailureTokenIntrospectionTimeoutCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginFailureTokenIntrospectionTimeoutCount

`func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureTokenIntrospectionTimeoutCount(v int64)`

SetLoginFailureTokenIntrospectionTimeoutCount sets LoginFailureTokenIntrospectionTimeoutCount field to given value.

### HasLoginFailureTokenIntrospectionTimeoutCount

`func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureTokenIntrospectionTimeoutCount() bool`

HasLoginFailureTokenIntrospectionTimeoutCount returns a boolean if a field has been set.

### GetLoginFailureTokenNotValidYetCount

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureTokenNotValidYetCount() int64`

GetLoginFailureTokenNotValidYetCount returns the LoginFailureTokenNotValidYetCount field if non-nil, zero value otherwise.

### GetLoginFailureTokenNotValidYetCountOk

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureTokenNotValidYetCountOk() (*int64, bool)`

GetLoginFailureTokenNotValidYetCountOk returns a tuple with the LoginFailureTokenNotValidYetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginFailureTokenNotValidYetCount

`func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureTokenNotValidYetCount(v int64)`

SetLoginFailureTokenNotValidYetCount sets LoginFailureTokenNotValidYetCount field to given value.

### HasLoginFailureTokenNotValidYetCount

`func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureTokenNotValidYetCount() bool`

HasLoginFailureTokenNotValidYetCount returns a boolean if a field has been set.

### GetLoginFailureUnsupportedAlgCount

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureUnsupportedAlgCount() int64`

GetLoginFailureUnsupportedAlgCount returns the LoginFailureUnsupportedAlgCount field if non-nil, zero value otherwise.

### GetLoginFailureUnsupportedAlgCountOk

`func (o *MsgVpnAuthenticationOauthProvider) GetLoginFailureUnsupportedAlgCountOk() (*int64, bool)`

GetLoginFailureUnsupportedAlgCountOk returns a tuple with the LoginFailureUnsupportedAlgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginFailureUnsupportedAlgCount

`func (o *MsgVpnAuthenticationOauthProvider) SetLoginFailureUnsupportedAlgCount(v int64)`

SetLoginFailureUnsupportedAlgCount sets LoginFailureUnsupportedAlgCount field to given value.

### HasLoginFailureUnsupportedAlgCount

`func (o *MsgVpnAuthenticationOauthProvider) HasLoginFailureUnsupportedAlgCount() bool`

HasLoginFailureUnsupportedAlgCount returns a boolean if a field has been set.

### GetMissingAuthorizationGroupCount

`func (o *MsgVpnAuthenticationOauthProvider) GetMissingAuthorizationGroupCount() int64`

GetMissingAuthorizationGroupCount returns the MissingAuthorizationGroupCount field if non-nil, zero value otherwise.

### GetMissingAuthorizationGroupCountOk

`func (o *MsgVpnAuthenticationOauthProvider) GetMissingAuthorizationGroupCountOk() (*int64, bool)`

GetMissingAuthorizationGroupCountOk returns a tuple with the MissingAuthorizationGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingAuthorizationGroupCount

`func (o *MsgVpnAuthenticationOauthProvider) SetMissingAuthorizationGroupCount(v int64)`

SetMissingAuthorizationGroupCount sets MissingAuthorizationGroupCount field to given value.

### HasMissingAuthorizationGroupCount

`func (o *MsgVpnAuthenticationOauthProvider) HasMissingAuthorizationGroupCount() bool`

HasMissingAuthorizationGroupCount returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnAuthenticationOauthProvider) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnAuthenticationOauthProvider) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnAuthenticationOauthProvider) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnAuthenticationOauthProvider) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetOauthProviderName

`func (o *MsgVpnAuthenticationOauthProvider) GetOauthProviderName() string`

GetOauthProviderName returns the OauthProviderName field if non-nil, zero value otherwise.

### GetOauthProviderNameOk

`func (o *MsgVpnAuthenticationOauthProvider) GetOauthProviderNameOk() (*string, bool)`

GetOauthProviderNameOk returns a tuple with the OauthProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthProviderName

`func (o *MsgVpnAuthenticationOauthProvider) SetOauthProviderName(v string)`

SetOauthProviderName sets OauthProviderName field to given value.

### HasOauthProviderName

`func (o *MsgVpnAuthenticationOauthProvider) HasOauthProviderName() bool`

HasOauthProviderName returns a boolean if a field has been set.

### GetTokenIgnoreTimeLimitsEnabled

`func (o *MsgVpnAuthenticationOauthProvider) GetTokenIgnoreTimeLimitsEnabled() bool`

GetTokenIgnoreTimeLimitsEnabled returns the TokenIgnoreTimeLimitsEnabled field if non-nil, zero value otherwise.

### GetTokenIgnoreTimeLimitsEnabledOk

`func (o *MsgVpnAuthenticationOauthProvider) GetTokenIgnoreTimeLimitsEnabledOk() (*bool, bool)`

GetTokenIgnoreTimeLimitsEnabledOk returns a tuple with the TokenIgnoreTimeLimitsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIgnoreTimeLimitsEnabled

`func (o *MsgVpnAuthenticationOauthProvider) SetTokenIgnoreTimeLimitsEnabled(v bool)`

SetTokenIgnoreTimeLimitsEnabled sets TokenIgnoreTimeLimitsEnabled field to given value.

### HasTokenIgnoreTimeLimitsEnabled

`func (o *MsgVpnAuthenticationOauthProvider) HasTokenIgnoreTimeLimitsEnabled() bool`

HasTokenIgnoreTimeLimitsEnabled returns a boolean if a field has been set.

### GetTokenIntrospectionAverageTime

`func (o *MsgVpnAuthenticationOauthProvider) GetTokenIntrospectionAverageTime() int32`

GetTokenIntrospectionAverageTime returns the TokenIntrospectionAverageTime field if non-nil, zero value otherwise.

### GetTokenIntrospectionAverageTimeOk

`func (o *MsgVpnAuthenticationOauthProvider) GetTokenIntrospectionAverageTimeOk() (*int32, bool)`

GetTokenIntrospectionAverageTimeOk returns a tuple with the TokenIntrospectionAverageTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIntrospectionAverageTime

`func (o *MsgVpnAuthenticationOauthProvider) SetTokenIntrospectionAverageTime(v int32)`

SetTokenIntrospectionAverageTime sets TokenIntrospectionAverageTime field to given value.

### HasTokenIntrospectionAverageTime

`func (o *MsgVpnAuthenticationOauthProvider) HasTokenIntrospectionAverageTime() bool`

HasTokenIntrospectionAverageTime returns a boolean if a field has been set.

### GetTokenIntrospectionLastFailureReason

`func (o *MsgVpnAuthenticationOauthProvider) GetTokenIntrospectionLastFailureReason() string`

GetTokenIntrospectionLastFailureReason returns the TokenIntrospectionLastFailureReason field if non-nil, zero value otherwise.

### GetTokenIntrospectionLastFailureReasonOk

`func (o *MsgVpnAuthenticationOauthProvider) GetTokenIntrospectionLastFailureReasonOk() (*string, bool)`

GetTokenIntrospectionLastFailureReasonOk returns a tuple with the TokenIntrospectionLastFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIntrospectionLastFailureReason

`func (o *MsgVpnAuthenticationOauthProvider) SetTokenIntrospectionLastFailureReason(v string)`

SetTokenIntrospectionLastFailureReason sets TokenIntrospectionLastFailureReason field to given value.

### HasTokenIntrospectionLastFailureReason

`func (o *MsgVpnAuthenticationOauthProvider) HasTokenIntrospectionLastFailureReason() bool`

HasTokenIntrospectionLastFailureReason returns a boolean if a field has been set.

### GetTokenIntrospectionLastFailureTime

`func (o *MsgVpnAuthenticationOauthProvider) GetTokenIntrospectionLastFailureTime() int32`

GetTokenIntrospectionLastFailureTime returns the TokenIntrospectionLastFailureTime field if non-nil, zero value otherwise.

### GetTokenIntrospectionLastFailureTimeOk

`func (o *MsgVpnAuthenticationOauthProvider) GetTokenIntrospectionLastFailureTimeOk() (*int32, bool)`

GetTokenIntrospectionLastFailureTimeOk returns a tuple with the TokenIntrospectionLastFailureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIntrospectionLastFailureTime

`func (o *MsgVpnAuthenticationOauthProvider) SetTokenIntrospectionLastFailureTime(v int32)`

SetTokenIntrospectionLastFailureTime sets TokenIntrospectionLastFailureTime field to given value.

### HasTokenIntrospectionLastFailureTime

`func (o *MsgVpnAuthenticationOauthProvider) HasTokenIntrospectionLastFailureTime() bool`

HasTokenIntrospectionLastFailureTime returns a boolean if a field has been set.

### GetTokenIntrospectionParameterName

`func (o *MsgVpnAuthenticationOauthProvider) GetTokenIntrospectionParameterName() string`

GetTokenIntrospectionParameterName returns the TokenIntrospectionParameterName field if non-nil, zero value otherwise.

### GetTokenIntrospectionParameterNameOk

`func (o *MsgVpnAuthenticationOauthProvider) GetTokenIntrospectionParameterNameOk() (*string, bool)`

GetTokenIntrospectionParameterNameOk returns a tuple with the TokenIntrospectionParameterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIntrospectionParameterName

`func (o *MsgVpnAuthenticationOauthProvider) SetTokenIntrospectionParameterName(v string)`

SetTokenIntrospectionParameterName sets TokenIntrospectionParameterName field to given value.

### HasTokenIntrospectionParameterName

`func (o *MsgVpnAuthenticationOauthProvider) HasTokenIntrospectionParameterName() bool`

HasTokenIntrospectionParameterName returns a boolean if a field has been set.

### GetTokenIntrospectionSuccessCount

`func (o *MsgVpnAuthenticationOauthProvider) GetTokenIntrospectionSuccessCount() int64`

GetTokenIntrospectionSuccessCount returns the TokenIntrospectionSuccessCount field if non-nil, zero value otherwise.

### GetTokenIntrospectionSuccessCountOk

`func (o *MsgVpnAuthenticationOauthProvider) GetTokenIntrospectionSuccessCountOk() (*int64, bool)`

GetTokenIntrospectionSuccessCountOk returns a tuple with the TokenIntrospectionSuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIntrospectionSuccessCount

`func (o *MsgVpnAuthenticationOauthProvider) SetTokenIntrospectionSuccessCount(v int64)`

SetTokenIntrospectionSuccessCount sets TokenIntrospectionSuccessCount field to given value.

### HasTokenIntrospectionSuccessCount

`func (o *MsgVpnAuthenticationOauthProvider) HasTokenIntrospectionSuccessCount() bool`

HasTokenIntrospectionSuccessCount returns a boolean if a field has been set.

### GetTokenIntrospectionTimeout

`func (o *MsgVpnAuthenticationOauthProvider) GetTokenIntrospectionTimeout() int32`

GetTokenIntrospectionTimeout returns the TokenIntrospectionTimeout field if non-nil, zero value otherwise.

### GetTokenIntrospectionTimeoutOk

`func (o *MsgVpnAuthenticationOauthProvider) GetTokenIntrospectionTimeoutOk() (*int32, bool)`

GetTokenIntrospectionTimeoutOk returns a tuple with the TokenIntrospectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIntrospectionTimeout

`func (o *MsgVpnAuthenticationOauthProvider) SetTokenIntrospectionTimeout(v int32)`

SetTokenIntrospectionTimeout sets TokenIntrospectionTimeout field to given value.

### HasTokenIntrospectionTimeout

`func (o *MsgVpnAuthenticationOauthProvider) HasTokenIntrospectionTimeout() bool`

HasTokenIntrospectionTimeout returns a boolean if a field has been set.

### GetTokenIntrospectionUri

`func (o *MsgVpnAuthenticationOauthProvider) GetTokenIntrospectionUri() string`

GetTokenIntrospectionUri returns the TokenIntrospectionUri field if non-nil, zero value otherwise.

### GetTokenIntrospectionUriOk

`func (o *MsgVpnAuthenticationOauthProvider) GetTokenIntrospectionUriOk() (*string, bool)`

GetTokenIntrospectionUriOk returns a tuple with the TokenIntrospectionUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIntrospectionUri

`func (o *MsgVpnAuthenticationOauthProvider) SetTokenIntrospectionUri(v string)`

SetTokenIntrospectionUri sets TokenIntrospectionUri field to given value.

### HasTokenIntrospectionUri

`func (o *MsgVpnAuthenticationOauthProvider) HasTokenIntrospectionUri() bool`

HasTokenIntrospectionUri returns a boolean if a field has been set.

### GetTokenIntrospectionUsername

`func (o *MsgVpnAuthenticationOauthProvider) GetTokenIntrospectionUsername() string`

GetTokenIntrospectionUsername returns the TokenIntrospectionUsername field if non-nil, zero value otherwise.

### GetTokenIntrospectionUsernameOk

`func (o *MsgVpnAuthenticationOauthProvider) GetTokenIntrospectionUsernameOk() (*string, bool)`

GetTokenIntrospectionUsernameOk returns a tuple with the TokenIntrospectionUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIntrospectionUsername

`func (o *MsgVpnAuthenticationOauthProvider) SetTokenIntrospectionUsername(v string)`

SetTokenIntrospectionUsername sets TokenIntrospectionUsername field to given value.

### HasTokenIntrospectionUsername

`func (o *MsgVpnAuthenticationOauthProvider) HasTokenIntrospectionUsername() bool`

HasTokenIntrospectionUsername returns a boolean if a field has been set.

### GetUsernameClaimName

`func (o *MsgVpnAuthenticationOauthProvider) GetUsernameClaimName() string`

GetUsernameClaimName returns the UsernameClaimName field if non-nil, zero value otherwise.

### GetUsernameClaimNameOk

`func (o *MsgVpnAuthenticationOauthProvider) GetUsernameClaimNameOk() (*string, bool)`

GetUsernameClaimNameOk returns a tuple with the UsernameClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameClaimName

`func (o *MsgVpnAuthenticationOauthProvider) SetUsernameClaimName(v string)`

SetUsernameClaimName sets UsernameClaimName field to given value.

### HasUsernameClaimName

`func (o *MsgVpnAuthenticationOauthProvider) HasUsernameClaimName() bool`

HasUsernameClaimName returns a boolean if a field has been set.

### GetUsernameClaimSource

`func (o *MsgVpnAuthenticationOauthProvider) GetUsernameClaimSource() string`

GetUsernameClaimSource returns the UsernameClaimSource field if non-nil, zero value otherwise.

### GetUsernameClaimSourceOk

`func (o *MsgVpnAuthenticationOauthProvider) GetUsernameClaimSourceOk() (*string, bool)`

GetUsernameClaimSourceOk returns a tuple with the UsernameClaimSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameClaimSource

`func (o *MsgVpnAuthenticationOauthProvider) SetUsernameClaimSource(v string)`

SetUsernameClaimSource sets UsernameClaimSource field to given value.

### HasUsernameClaimSource

`func (o *MsgVpnAuthenticationOauthProvider) HasUsernameClaimSource() bool`

HasUsernameClaimSource returns a boolean if a field has been set.

### GetUsernameValidateEnabled

`func (o *MsgVpnAuthenticationOauthProvider) GetUsernameValidateEnabled() bool`

GetUsernameValidateEnabled returns the UsernameValidateEnabled field if non-nil, zero value otherwise.

### GetUsernameValidateEnabledOk

`func (o *MsgVpnAuthenticationOauthProvider) GetUsernameValidateEnabledOk() (*bool, bool)`

GetUsernameValidateEnabledOk returns a tuple with the UsernameValidateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameValidateEnabled

`func (o *MsgVpnAuthenticationOauthProvider) SetUsernameValidateEnabled(v bool)`

SetUsernameValidateEnabled sets UsernameValidateEnabled field to given value.

### HasUsernameValidateEnabled

`func (o *MsgVpnAuthenticationOauthProvider) HasUsernameValidateEnabled() bool`

HasUsernameValidateEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


