# MsgVpnBridge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BridgeName** | Pointer to **string** | The name of the Bridge. | [optional] 
**BridgeVirtualRouter** | Pointer to **string** | The virtual router of the Bridge. The allowed values and their meaning are:  &lt;pre&gt; \&quot;primary\&quot; - The Bridge is used for the primary virtual router. \&quot;backup\&quot; - The Bridge is used for the backup virtual router. \&quot;auto\&quot; - The Bridge is automatically assigned a virtual router at creation, depending on the broker&#39;s active-standby role. &lt;/pre&gt;  | [optional] 
**Enabled** | Pointer to **bool** | Enable or disable the Bridge. The default value is &#x60;false&#x60;. | [optional] 
**MaxTtl** | Pointer to **int64** | The maximum time-to-live (TTL) in hops. Messages are discarded if their TTL exceeds this value. The default value is &#x60;8&#x60;. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**RemoteAuthenticationBasicClientUsername** | Pointer to **string** | The Client Username the Bridge uses to login to the remote Message VPN. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] 
**RemoteAuthenticationBasicPassword** | Pointer to **string** | The password for the Client Username. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] 
**RemoteAuthenticationClientCertContent** | Pointer to **string** | The PEM formatted content for the client certificate used by the Bridge to login to the remote Message VPN. It must consist of a private key and between one and three certificates comprising the certificate trust chain. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. Changing this attribute requires an HTTPS connection. The default value is &#x60;\&quot;\&quot;&#x60;. Available since 2.9. | [optional] 
**RemoteAuthenticationClientCertPassword** | Pointer to **string** | The password for the client certificate. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. Changing this attribute requires an HTTPS connection. The default value is &#x60;\&quot;\&quot;&#x60;. Available since 2.9. | [optional] 
**RemoteAuthenticationScheme** | Pointer to **string** | The authentication scheme for the remote Message VPN. The default value is &#x60;\&quot;basic\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;basic\&quot; - Basic Authentication Scheme (via username and password). \&quot;client-certificate\&quot; - Client Certificate Authentication Scheme (via certificate file or content). &lt;/pre&gt;  | [optional] 
**RemoteConnectionRetryCount** | Pointer to **int64** | The maximum number of retry attempts to establish a connection to the remote Message VPN. A value of 0 means to retry forever. The default value is &#x60;0&#x60;. | [optional] 
**RemoteConnectionRetryDelay** | Pointer to **int64** | The number of seconds the broker waits for the bridge connection to be established before attempting a new connection. The default value is &#x60;3&#x60;. | [optional] 
**RemoteDeliverToOnePriority** | Pointer to **string** | The priority for deliver-to-one (DTO) messages transmitted from the remote Message VPN. The default value is &#x60;\&quot;p1\&quot;&#x60;. The allowed values and their meaning are:  &lt;pre&gt; \&quot;p1\&quot; - The 1st or highest priority. \&quot;p2\&quot; - The 2nd highest priority. \&quot;p3\&quot; - The 3rd highest priority. \&quot;p4\&quot; - The 4th highest priority. \&quot;da\&quot; - Ignore priority and deliver always. &lt;/pre&gt;  | [optional] 
**TlsCipherSuiteList** | Pointer to **string** | The colon-separated list of cipher suites supported for TLS connections to the remote Message VPN. The value \&quot;default\&quot; implies all supported suites ordered from most secure to least secure. The default value is &#x60;\&quot;default\&quot;&#x60;. | [optional] 

## Methods

### NewMsgVpnBridge

`func NewMsgVpnBridge() *MsgVpnBridge`

NewMsgVpnBridge instantiates a new MsgVpnBridge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnBridgeWithDefaults

`func NewMsgVpnBridgeWithDefaults() *MsgVpnBridge`

NewMsgVpnBridgeWithDefaults instantiates a new MsgVpnBridge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBridgeName

`func (o *MsgVpnBridge) GetBridgeName() string`

GetBridgeName returns the BridgeName field if non-nil, zero value otherwise.

### GetBridgeNameOk

`func (o *MsgVpnBridge) GetBridgeNameOk() (*string, bool)`

GetBridgeNameOk returns a tuple with the BridgeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeName

`func (o *MsgVpnBridge) SetBridgeName(v string)`

SetBridgeName sets BridgeName field to given value.

### HasBridgeName

`func (o *MsgVpnBridge) HasBridgeName() bool`

HasBridgeName returns a boolean if a field has been set.

### GetBridgeVirtualRouter

`func (o *MsgVpnBridge) GetBridgeVirtualRouter() string`

GetBridgeVirtualRouter returns the BridgeVirtualRouter field if non-nil, zero value otherwise.

### GetBridgeVirtualRouterOk

`func (o *MsgVpnBridge) GetBridgeVirtualRouterOk() (*string, bool)`

GetBridgeVirtualRouterOk returns a tuple with the BridgeVirtualRouter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeVirtualRouter

`func (o *MsgVpnBridge) SetBridgeVirtualRouter(v string)`

SetBridgeVirtualRouter sets BridgeVirtualRouter field to given value.

### HasBridgeVirtualRouter

`func (o *MsgVpnBridge) HasBridgeVirtualRouter() bool`

HasBridgeVirtualRouter returns a boolean if a field has been set.

### GetEnabled

`func (o *MsgVpnBridge) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MsgVpnBridge) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MsgVpnBridge) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MsgVpnBridge) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMaxTtl

`func (o *MsgVpnBridge) GetMaxTtl() int64`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *MsgVpnBridge) GetMaxTtlOk() (*int64, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *MsgVpnBridge) SetMaxTtl(v int64)`

SetMaxTtl sets MaxTtl field to given value.

### HasMaxTtl

`func (o *MsgVpnBridge) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnBridge) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnBridge) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnBridge) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnBridge) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetRemoteAuthenticationBasicClientUsername

`func (o *MsgVpnBridge) GetRemoteAuthenticationBasicClientUsername() string`

GetRemoteAuthenticationBasicClientUsername returns the RemoteAuthenticationBasicClientUsername field if non-nil, zero value otherwise.

### GetRemoteAuthenticationBasicClientUsernameOk

`func (o *MsgVpnBridge) GetRemoteAuthenticationBasicClientUsernameOk() (*string, bool)`

GetRemoteAuthenticationBasicClientUsernameOk returns a tuple with the RemoteAuthenticationBasicClientUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAuthenticationBasicClientUsername

`func (o *MsgVpnBridge) SetRemoteAuthenticationBasicClientUsername(v string)`

SetRemoteAuthenticationBasicClientUsername sets RemoteAuthenticationBasicClientUsername field to given value.

### HasRemoteAuthenticationBasicClientUsername

`func (o *MsgVpnBridge) HasRemoteAuthenticationBasicClientUsername() bool`

HasRemoteAuthenticationBasicClientUsername returns a boolean if a field has been set.

### GetRemoteAuthenticationBasicPassword

`func (o *MsgVpnBridge) GetRemoteAuthenticationBasicPassword() string`

GetRemoteAuthenticationBasicPassword returns the RemoteAuthenticationBasicPassword field if non-nil, zero value otherwise.

### GetRemoteAuthenticationBasicPasswordOk

`func (o *MsgVpnBridge) GetRemoteAuthenticationBasicPasswordOk() (*string, bool)`

GetRemoteAuthenticationBasicPasswordOk returns a tuple with the RemoteAuthenticationBasicPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAuthenticationBasicPassword

`func (o *MsgVpnBridge) SetRemoteAuthenticationBasicPassword(v string)`

SetRemoteAuthenticationBasicPassword sets RemoteAuthenticationBasicPassword field to given value.

### HasRemoteAuthenticationBasicPassword

`func (o *MsgVpnBridge) HasRemoteAuthenticationBasicPassword() bool`

HasRemoteAuthenticationBasicPassword returns a boolean if a field has been set.

### GetRemoteAuthenticationClientCertContent

`func (o *MsgVpnBridge) GetRemoteAuthenticationClientCertContent() string`

GetRemoteAuthenticationClientCertContent returns the RemoteAuthenticationClientCertContent field if non-nil, zero value otherwise.

### GetRemoteAuthenticationClientCertContentOk

`func (o *MsgVpnBridge) GetRemoteAuthenticationClientCertContentOk() (*string, bool)`

GetRemoteAuthenticationClientCertContentOk returns a tuple with the RemoteAuthenticationClientCertContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAuthenticationClientCertContent

`func (o *MsgVpnBridge) SetRemoteAuthenticationClientCertContent(v string)`

SetRemoteAuthenticationClientCertContent sets RemoteAuthenticationClientCertContent field to given value.

### HasRemoteAuthenticationClientCertContent

`func (o *MsgVpnBridge) HasRemoteAuthenticationClientCertContent() bool`

HasRemoteAuthenticationClientCertContent returns a boolean if a field has been set.

### GetRemoteAuthenticationClientCertPassword

`func (o *MsgVpnBridge) GetRemoteAuthenticationClientCertPassword() string`

GetRemoteAuthenticationClientCertPassword returns the RemoteAuthenticationClientCertPassword field if non-nil, zero value otherwise.

### GetRemoteAuthenticationClientCertPasswordOk

`func (o *MsgVpnBridge) GetRemoteAuthenticationClientCertPasswordOk() (*string, bool)`

GetRemoteAuthenticationClientCertPasswordOk returns a tuple with the RemoteAuthenticationClientCertPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAuthenticationClientCertPassword

`func (o *MsgVpnBridge) SetRemoteAuthenticationClientCertPassword(v string)`

SetRemoteAuthenticationClientCertPassword sets RemoteAuthenticationClientCertPassword field to given value.

### HasRemoteAuthenticationClientCertPassword

`func (o *MsgVpnBridge) HasRemoteAuthenticationClientCertPassword() bool`

HasRemoteAuthenticationClientCertPassword returns a boolean if a field has been set.

### GetRemoteAuthenticationScheme

`func (o *MsgVpnBridge) GetRemoteAuthenticationScheme() string`

GetRemoteAuthenticationScheme returns the RemoteAuthenticationScheme field if non-nil, zero value otherwise.

### GetRemoteAuthenticationSchemeOk

`func (o *MsgVpnBridge) GetRemoteAuthenticationSchemeOk() (*string, bool)`

GetRemoteAuthenticationSchemeOk returns a tuple with the RemoteAuthenticationScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAuthenticationScheme

`func (o *MsgVpnBridge) SetRemoteAuthenticationScheme(v string)`

SetRemoteAuthenticationScheme sets RemoteAuthenticationScheme field to given value.

### HasRemoteAuthenticationScheme

`func (o *MsgVpnBridge) HasRemoteAuthenticationScheme() bool`

HasRemoteAuthenticationScheme returns a boolean if a field has been set.

### GetRemoteConnectionRetryCount

`func (o *MsgVpnBridge) GetRemoteConnectionRetryCount() int64`

GetRemoteConnectionRetryCount returns the RemoteConnectionRetryCount field if non-nil, zero value otherwise.

### GetRemoteConnectionRetryCountOk

`func (o *MsgVpnBridge) GetRemoteConnectionRetryCountOk() (*int64, bool)`

GetRemoteConnectionRetryCountOk returns a tuple with the RemoteConnectionRetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteConnectionRetryCount

`func (o *MsgVpnBridge) SetRemoteConnectionRetryCount(v int64)`

SetRemoteConnectionRetryCount sets RemoteConnectionRetryCount field to given value.

### HasRemoteConnectionRetryCount

`func (o *MsgVpnBridge) HasRemoteConnectionRetryCount() bool`

HasRemoteConnectionRetryCount returns a boolean if a field has been set.

### GetRemoteConnectionRetryDelay

`func (o *MsgVpnBridge) GetRemoteConnectionRetryDelay() int64`

GetRemoteConnectionRetryDelay returns the RemoteConnectionRetryDelay field if non-nil, zero value otherwise.

### GetRemoteConnectionRetryDelayOk

`func (o *MsgVpnBridge) GetRemoteConnectionRetryDelayOk() (*int64, bool)`

GetRemoteConnectionRetryDelayOk returns a tuple with the RemoteConnectionRetryDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteConnectionRetryDelay

`func (o *MsgVpnBridge) SetRemoteConnectionRetryDelay(v int64)`

SetRemoteConnectionRetryDelay sets RemoteConnectionRetryDelay field to given value.

### HasRemoteConnectionRetryDelay

`func (o *MsgVpnBridge) HasRemoteConnectionRetryDelay() bool`

HasRemoteConnectionRetryDelay returns a boolean if a field has been set.

### GetRemoteDeliverToOnePriority

`func (o *MsgVpnBridge) GetRemoteDeliverToOnePriority() string`

GetRemoteDeliverToOnePriority returns the RemoteDeliverToOnePriority field if non-nil, zero value otherwise.

### GetRemoteDeliverToOnePriorityOk

`func (o *MsgVpnBridge) GetRemoteDeliverToOnePriorityOk() (*string, bool)`

GetRemoteDeliverToOnePriorityOk returns a tuple with the RemoteDeliverToOnePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteDeliverToOnePriority

`func (o *MsgVpnBridge) SetRemoteDeliverToOnePriority(v string)`

SetRemoteDeliverToOnePriority sets RemoteDeliverToOnePriority field to given value.

### HasRemoteDeliverToOnePriority

`func (o *MsgVpnBridge) HasRemoteDeliverToOnePriority() bool`

HasRemoteDeliverToOnePriority returns a boolean if a field has been set.

### GetTlsCipherSuiteList

`func (o *MsgVpnBridge) GetTlsCipherSuiteList() string`

GetTlsCipherSuiteList returns the TlsCipherSuiteList field if non-nil, zero value otherwise.

### GetTlsCipherSuiteListOk

`func (o *MsgVpnBridge) GetTlsCipherSuiteListOk() (*string, bool)`

GetTlsCipherSuiteListOk returns a tuple with the TlsCipherSuiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCipherSuiteList

`func (o *MsgVpnBridge) SetTlsCipherSuiteList(v string)`

SetTlsCipherSuiteList sets TlsCipherSuiteList field to given value.

### HasTlsCipherSuiteList

`func (o *MsgVpnBridge) HasTlsCipherSuiteList() bool`

HasTlsCipherSuiteList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


