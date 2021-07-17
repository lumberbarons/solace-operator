# MsgVpnBridgeLocalSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BridgeName** | Pointer to **string** | The name of the Bridge. | [optional] 
**BridgeVirtualRouter** | Pointer to **string** | The virtual router of the Bridge. The allowed values and their meaning are:  &lt;pre&gt; \&quot;primary\&quot; - The Bridge is used for the primary virtual router. \&quot;backup\&quot; - The Bridge is used for the backup virtual router. \&quot;auto\&quot; - The Bridge is automatically assigned a virtual router at creation, depending on the broker&#39;s active-standby role. &lt;/pre&gt;  | [optional] 
**DtoPriority** | Pointer to **string** | The priority of the Bridge local subscription topic for receiving deliver-to-one (DTO) messages. The allowed values and their meaning are:  &lt;pre&gt; \&quot;p1\&quot; - The 1st or highest priority. \&quot;p2\&quot; - The 2nd highest priority. \&quot;p3\&quot; - The 3rd highest priority. \&quot;p4\&quot; - The 4th highest priority. \&quot;da\&quot; - Ignore priority and deliver always. &lt;/pre&gt;  | [optional] 
**LocalSubscriptionTopic** | Pointer to **string** | The topic of the Bridge local subscription. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 

## Methods

### NewMsgVpnBridgeLocalSubscription

`func NewMsgVpnBridgeLocalSubscription() *MsgVpnBridgeLocalSubscription`

NewMsgVpnBridgeLocalSubscription instantiates a new MsgVpnBridgeLocalSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnBridgeLocalSubscriptionWithDefaults

`func NewMsgVpnBridgeLocalSubscriptionWithDefaults() *MsgVpnBridgeLocalSubscription`

NewMsgVpnBridgeLocalSubscriptionWithDefaults instantiates a new MsgVpnBridgeLocalSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBridgeName

`func (o *MsgVpnBridgeLocalSubscription) GetBridgeName() string`

GetBridgeName returns the BridgeName field if non-nil, zero value otherwise.

### GetBridgeNameOk

`func (o *MsgVpnBridgeLocalSubscription) GetBridgeNameOk() (*string, bool)`

GetBridgeNameOk returns a tuple with the BridgeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeName

`func (o *MsgVpnBridgeLocalSubscription) SetBridgeName(v string)`

SetBridgeName sets BridgeName field to given value.

### HasBridgeName

`func (o *MsgVpnBridgeLocalSubscription) HasBridgeName() bool`

HasBridgeName returns a boolean if a field has been set.

### GetBridgeVirtualRouter

`func (o *MsgVpnBridgeLocalSubscription) GetBridgeVirtualRouter() string`

GetBridgeVirtualRouter returns the BridgeVirtualRouter field if non-nil, zero value otherwise.

### GetBridgeVirtualRouterOk

`func (o *MsgVpnBridgeLocalSubscription) GetBridgeVirtualRouterOk() (*string, bool)`

GetBridgeVirtualRouterOk returns a tuple with the BridgeVirtualRouter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeVirtualRouter

`func (o *MsgVpnBridgeLocalSubscription) SetBridgeVirtualRouter(v string)`

SetBridgeVirtualRouter sets BridgeVirtualRouter field to given value.

### HasBridgeVirtualRouter

`func (o *MsgVpnBridgeLocalSubscription) HasBridgeVirtualRouter() bool`

HasBridgeVirtualRouter returns a boolean if a field has been set.

### GetDtoPriority

`func (o *MsgVpnBridgeLocalSubscription) GetDtoPriority() string`

GetDtoPriority returns the DtoPriority field if non-nil, zero value otherwise.

### GetDtoPriorityOk

`func (o *MsgVpnBridgeLocalSubscription) GetDtoPriorityOk() (*string, bool)`

GetDtoPriorityOk returns a tuple with the DtoPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtoPriority

`func (o *MsgVpnBridgeLocalSubscription) SetDtoPriority(v string)`

SetDtoPriority sets DtoPriority field to given value.

### HasDtoPriority

`func (o *MsgVpnBridgeLocalSubscription) HasDtoPriority() bool`

HasDtoPriority returns a boolean if a field has been set.

### GetLocalSubscriptionTopic

`func (o *MsgVpnBridgeLocalSubscription) GetLocalSubscriptionTopic() string`

GetLocalSubscriptionTopic returns the LocalSubscriptionTopic field if non-nil, zero value otherwise.

### GetLocalSubscriptionTopicOk

`func (o *MsgVpnBridgeLocalSubscription) GetLocalSubscriptionTopicOk() (*string, bool)`

GetLocalSubscriptionTopicOk returns a tuple with the LocalSubscriptionTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalSubscriptionTopic

`func (o *MsgVpnBridgeLocalSubscription) SetLocalSubscriptionTopic(v string)`

SetLocalSubscriptionTopic sets LocalSubscriptionTopic field to given value.

### HasLocalSubscriptionTopic

`func (o *MsgVpnBridgeLocalSubscription) HasLocalSubscriptionTopic() bool`

HasLocalSubscriptionTopic returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnBridgeLocalSubscription) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnBridgeLocalSubscription) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnBridgeLocalSubscription) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnBridgeLocalSubscription) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


