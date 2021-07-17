# MsgVpnBridgeTlsTrustedCommonName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BridgeName** | Pointer to **string** | The name of the Bridge. Deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation. | [optional] 
**BridgeVirtualRouter** | Pointer to **string** | The virtual router of the Bridge. The allowed values and their meaning are:  &lt;pre&gt; \&quot;primary\&quot; - The Bridge is used for the primary virtual router. \&quot;backup\&quot; - The Bridge is used for the backup virtual router. \&quot;auto\&quot; - The Bridge is automatically assigned a virtual router at creation, depending on the broker&#39;s active-standby role. &lt;/pre&gt;  Deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. Deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation. | [optional] 
**TlsTrustedCommonName** | Pointer to **string** | The expected trusted common name of the remote certificate. Deprecated since 2.18. Common Name validation has been replaced by Server Certificate Name validation. | [optional] 

## Methods

### NewMsgVpnBridgeTlsTrustedCommonName

`func NewMsgVpnBridgeTlsTrustedCommonName() *MsgVpnBridgeTlsTrustedCommonName`

NewMsgVpnBridgeTlsTrustedCommonName instantiates a new MsgVpnBridgeTlsTrustedCommonName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnBridgeTlsTrustedCommonNameWithDefaults

`func NewMsgVpnBridgeTlsTrustedCommonNameWithDefaults() *MsgVpnBridgeTlsTrustedCommonName`

NewMsgVpnBridgeTlsTrustedCommonNameWithDefaults instantiates a new MsgVpnBridgeTlsTrustedCommonName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBridgeName

`func (o *MsgVpnBridgeTlsTrustedCommonName) GetBridgeName() string`

GetBridgeName returns the BridgeName field if non-nil, zero value otherwise.

### GetBridgeNameOk

`func (o *MsgVpnBridgeTlsTrustedCommonName) GetBridgeNameOk() (*string, bool)`

GetBridgeNameOk returns a tuple with the BridgeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeName

`func (o *MsgVpnBridgeTlsTrustedCommonName) SetBridgeName(v string)`

SetBridgeName sets BridgeName field to given value.

### HasBridgeName

`func (o *MsgVpnBridgeTlsTrustedCommonName) HasBridgeName() bool`

HasBridgeName returns a boolean if a field has been set.

### GetBridgeVirtualRouter

`func (o *MsgVpnBridgeTlsTrustedCommonName) GetBridgeVirtualRouter() string`

GetBridgeVirtualRouter returns the BridgeVirtualRouter field if non-nil, zero value otherwise.

### GetBridgeVirtualRouterOk

`func (o *MsgVpnBridgeTlsTrustedCommonName) GetBridgeVirtualRouterOk() (*string, bool)`

GetBridgeVirtualRouterOk returns a tuple with the BridgeVirtualRouter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeVirtualRouter

`func (o *MsgVpnBridgeTlsTrustedCommonName) SetBridgeVirtualRouter(v string)`

SetBridgeVirtualRouter sets BridgeVirtualRouter field to given value.

### HasBridgeVirtualRouter

`func (o *MsgVpnBridgeTlsTrustedCommonName) HasBridgeVirtualRouter() bool`

HasBridgeVirtualRouter returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnBridgeTlsTrustedCommonName) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnBridgeTlsTrustedCommonName) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnBridgeTlsTrustedCommonName) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnBridgeTlsTrustedCommonName) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetTlsTrustedCommonName

`func (o *MsgVpnBridgeTlsTrustedCommonName) GetTlsTrustedCommonName() string`

GetTlsTrustedCommonName returns the TlsTrustedCommonName field if non-nil, zero value otherwise.

### GetTlsTrustedCommonNameOk

`func (o *MsgVpnBridgeTlsTrustedCommonName) GetTlsTrustedCommonNameOk() (*string, bool)`

GetTlsTrustedCommonNameOk returns a tuple with the TlsTrustedCommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsTrustedCommonName

`func (o *MsgVpnBridgeTlsTrustedCommonName) SetTlsTrustedCommonName(v string)`

SetTlsTrustedCommonName sets TlsTrustedCommonName field to given value.

### HasTlsTrustedCommonName

`func (o *MsgVpnBridgeTlsTrustedCommonName) HasTlsTrustedCommonName() bool`

HasTlsTrustedCommonName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


