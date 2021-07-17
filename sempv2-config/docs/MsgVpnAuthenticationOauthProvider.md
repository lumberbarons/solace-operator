# MsgVpnAuthenticationOauthProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudienceClaimName** | Pointer to **string** | The audience claim name, indicating which part of the object to use for determining the audience. The default value is &#x60;\&quot;aud\&quot;&#x60;. | [optional] 
**AudienceClaimSource** | Pointer to **string** | The audience claim source, indicating where to search for the audience value. The default value is &#x60;\&quot;id-token\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;access-token\&quot; - The OAuth v2 access_token. \&quot;id-token\&quot; - The OpenID Connect id_token. \&quot;introspection\&quot; - The result of introspecting the OAuth v2 access_token. &lt;/pre&gt;  | [optional] 
**AudienceClaimValue** | Pointer to **string** | The required audience value for a token to be considered valid. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] 
**AudienceValidationEnabled** | Pointer to **bool** | Enable or disable audience validation. The default value is &#x60;false&#x60;. | [optional] 
**AuthorizationGroupClaimName** | Pointer to **string** | The authorization group claim name, indicating which part of the object to use for determining the authorization group. The default value is &#x60;\&quot;scope\&quot;&#x60;. | [optional] 
**AuthorizationGroupClaimSource** | Pointer to **string** | The authorization group claim source, indicating where to search for the authorization group name. The default value is &#x60;\&quot;id-token\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;access-token\&quot; - The OAuth v2 access_token. \&quot;id-token\&quot; - The OpenID Connect id_token. \&quot;introspection\&quot; - The result of introspecting the OAuth v2 access_token. &lt;/pre&gt;  | [optional] 
**AuthorizationGroupEnabled** | Pointer to **bool** | Enable or disable OAuth based authorization. When enabled, the configured authorization type for OAuth clients is overridden. The default value is &#x60;false&#x60;. | [optional] 
**DisconnectOnTokenExpirationEnabled** | Pointer to **bool** | Enable or disable the disconnection of clients when their tokens expire. Changing this value does not affect existing clients, only new client connections. The default value is &#x60;true&#x60;. | [optional] 
**Enabled** | Pointer to **bool** | Enable or disable OAuth Provider client authentication. The default value is &#x60;false&#x60;. | [optional] 
**JwksRefreshInterval** | Pointer to **int32** | The number of seconds between forced JWKS public key refreshing. The default value is &#x60;86400&#x60;. | [optional] 
**JwksUri** | Pointer to **string** | The URI where the OAuth provider publishes its JWKS public keys. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**OauthProviderName** | Pointer to **string** | The name of the OAuth Provider. | [optional] 
**TokenIgnoreTimeLimitsEnabled** | Pointer to **bool** | Enable or disable whether to ignore time limits and accept tokens that are not yet valid or are no longer valid. The default value is &#x60;false&#x60;. | [optional] 
**TokenIntrospectionParameterName** | Pointer to **string** | The parameter name used to identify the token during access token introspection. A standards compliant OAuth introspection server expects \&quot;token\&quot;. The default value is &#x60;\&quot;token\&quot;&#x60;. | [optional] 
**TokenIntrospectionPassword** | Pointer to **string** | The password to use when logging into the token introspection URI. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] 
**TokenIntrospectionTimeout** | Pointer to **int32** | The maximum time in seconds a token introspection is allowed to take. The default value is &#x60;1&#x60;. | [optional] 
**TokenIntrospectionUri** | Pointer to **string** | The token introspection URI of the OAuth authentication server. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] 
**TokenIntrospectionUsername** | Pointer to **string** | The username to use when logging into the token introspection URI. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] 
**UsernameClaimName** | Pointer to **string** | The username claim name, indicating which part of the object to use for determining the username. The default value is &#x60;\&quot;sub\&quot;&#x60;. | [optional] 
**UsernameClaimSource** | Pointer to **string** | The username claim source, indicating where to search for the username value. The default value is &#x60;\&quot;id-token\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;access-token\&quot; - The OAuth v2 access_token. \&quot;id-token\&quot; - The OpenID Connect id_token. \&quot;introspection\&quot; - The result of introspecting the OAuth v2 access_token. &lt;/pre&gt;  | [optional] 
**UsernameValidateEnabled** | Pointer to **bool** | Enable or disable whether the API provided username will be validated against the username calculated from the token(s); the connection attempt is rejected if they differ. The default value is &#x60;false&#x60;. | [optional] 

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

### GetTokenIntrospectionPassword

`func (o *MsgVpnAuthenticationOauthProvider) GetTokenIntrospectionPassword() string`

GetTokenIntrospectionPassword returns the TokenIntrospectionPassword field if non-nil, zero value otherwise.

### GetTokenIntrospectionPasswordOk

`func (o *MsgVpnAuthenticationOauthProvider) GetTokenIntrospectionPasswordOk() (*string, bool)`

GetTokenIntrospectionPasswordOk returns a tuple with the TokenIntrospectionPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIntrospectionPassword

`func (o *MsgVpnAuthenticationOauthProvider) SetTokenIntrospectionPassword(v string)`

SetTokenIntrospectionPassword sets TokenIntrospectionPassword field to given value.

### HasTokenIntrospectionPassword

`func (o *MsgVpnAuthenticationOauthProvider) HasTokenIntrospectionPassword() bool`

HasTokenIntrospectionPassword returns a boolean if a field has been set.

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


