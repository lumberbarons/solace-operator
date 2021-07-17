# MsgVpnBridgeRemoteSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BridgeName** | Pointer to **string** | The name of the Bridge. | [optional] 
**BridgeVirtualRouter** | Pointer to **string** | The virtual router of the Bridge. The allowed values and their meaning are:  &lt;pre&gt; \&quot;primary\&quot; - The Bridge is used for the primary virtual router. \&quot;backup\&quot; - The Bridge is used for the backup virtual router. \&quot;auto\&quot; - The Bridge is automatically assigned a virtual router at creation, depending on the broker&#39;s active-standby role. &lt;/pre&gt;  | [optional] 
**DeliverAlwaysEnabled** | Pointer to **bool** | Indicates whether deliver-always is enabled for the Bridge remote subscription topic instead of a deliver-to-one remote priority. A given topic for the Bridge may be deliver-to-one or deliver-always but not both. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**RemoteSubscriptionTopic** | Pointer to **string** | The topic of the Bridge remote subscription. | [optional] 

## Methods

### NewMsgVpnBridgeRemoteSubscription

`func NewMsgVpnBridgeRemoteSubscription() *MsgVpnBridgeRemoteSubscription`

NewMsgVpnBridgeRemoteSubscription instantiates a new MsgVpnBridgeRemoteSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnBridgeRemoteSubscriptionWithDefaults

`func NewMsgVpnBridgeRemoteSubscriptionWithDefaults() *MsgVpnBridgeRemoteSubscription`

NewMsgVpnBridgeRemoteSubscriptionWithDefaults instantiates a new MsgVpnBridgeRemoteSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBridgeName

`func (o *MsgVpnBridgeRemoteSubscription) GetBridgeName() string`

GetBridgeName returns the BridgeName field if non-nil, zero value otherwise.

### GetBridgeNameOk

`func (o *MsgVpnBridgeRemoteSubscription) GetBridgeNameOk() (*string, bool)`

GetBridgeNameOk returns a tuple with the BridgeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeName

`func (o *MsgVpnBridgeRemoteSubscription) SetBridgeName(v string)`

SetBridgeName sets BridgeName field to given value.

### HasBridgeName

`func (o *MsgVpnBridgeRemoteSubscription) HasBridgeName() bool`

HasBridgeName returns a boolean if a field has been set.

### GetBridgeVirtualRouter

`func (o *MsgVpnBridgeRemoteSubscription) GetBridgeVirtualRouter() string`

GetBridgeVirtualRouter returns the BridgeVirtualRouter field if non-nil, zero value otherwise.

### GetBridgeVirtualRouterOk

`func (o *MsgVpnBridgeRemoteSubscription) GetBridgeVirtualRouterOk() (*string, bool)`

GetBridgeVirtualRouterOk returns a tuple with the BridgeVirtualRouter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeVirtualRouter

`func (o *MsgVpnBridgeRemoteSubscription) SetBridgeVirtualRouter(v string)`

SetBridgeVirtualRouter sets BridgeVirtualRouter field to given value.

### HasBridgeVirtualRouter

`func (o *MsgVpnBridgeRemoteSubscription) HasBridgeVirtualRouter() bool`

HasBridgeVirtualRouter returns a boolean if a field has been set.

### GetDeliverAlwaysEnabled

`func (o *MsgVpnBridgeRemoteSubscription) GetDeliverAlwaysEnabled() bool`

GetDeliverAlwaysEnabled returns the DeliverAlwaysEnabled field if non-nil, zero value otherwise.

### GetDeliverAlwaysEnabledOk

`func (o *MsgVpnBridgeRemoteSubscription) GetDeliverAlwaysEnabledOk() (*bool, bool)`

GetDeliverAlwaysEnabledOk returns a tuple with the DeliverAlwaysEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverAlwaysEnabled

`func (o *MsgVpnBridgeRemoteSubscription) SetDeliverAlwaysEnabled(v bool)`

SetDeliverAlwaysEnabled sets DeliverAlwaysEnabled field to given value.

### HasDeliverAlwaysEnabled

`func (o *MsgVpnBridgeRemoteSubscription) HasDeliverAlwaysEnabled() bool`

HasDeliverAlwaysEnabled returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnBridgeRemoteSubscription) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnBridgeRemoteSubscription) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnBridgeRemoteSubscription) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnBridgeRemoteSubscription) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetRemoteSubscriptionTopic

`func (o *MsgVpnBridgeRemoteSubscription) GetRemoteSubscriptionTopic() string`

GetRemoteSubscriptionTopic returns the RemoteSubscriptionTopic field if non-nil, zero value otherwise.

### GetRemoteSubscriptionTopicOk

`func (o *MsgVpnBridgeRemoteSubscription) GetRemoteSubscriptionTopicOk() (*string, bool)`

GetRemoteSubscriptionTopicOk returns a tuple with the RemoteSubscriptionTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSubscriptionTopic

`func (o *MsgVpnBridgeRemoteSubscription) SetRemoteSubscriptionTopic(v string)`

SetRemoteSubscriptionTopic sets RemoteSubscriptionTopic field to given value.

### HasRemoteSubscriptionTopic

`func (o *MsgVpnBridgeRemoteSubscription) HasRemoteSubscriptionTopic() bool`

HasRemoteSubscriptionTopic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


