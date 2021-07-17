# MsgVpnMqttSessionSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MqttSessionClientId** | Pointer to **string** | The Client ID of the MQTT Session, which corresponds to the ClientId provided in the MQTT CONNECT packet. | [optional] 
**MqttSessionVirtualRouter** | Pointer to **string** | The virtual router of the MQTT Session. The allowed values and their meaning are:  &lt;pre&gt; \&quot;primary\&quot; - The MQTT Session belongs to the primary virtual router. \&quot;backup\&quot; - The MQTT Session belongs to the backup virtual router. &lt;/pre&gt;  | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**SubscriptionQos** | Pointer to **int64** | The quality of service (QoS) for the MQTT Session subscription. | [optional] 
**SubscriptionTopic** | Pointer to **string** | The MQTT subscription topic. | [optional] 

## Methods

### NewMsgVpnMqttSessionSubscription

`func NewMsgVpnMqttSessionSubscription() *MsgVpnMqttSessionSubscription`

NewMsgVpnMqttSessionSubscription instantiates a new MsgVpnMqttSessionSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnMqttSessionSubscriptionWithDefaults

`func NewMsgVpnMqttSessionSubscriptionWithDefaults() *MsgVpnMqttSessionSubscription`

NewMsgVpnMqttSessionSubscriptionWithDefaults instantiates a new MsgVpnMqttSessionSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMqttSessionClientId

`func (o *MsgVpnMqttSessionSubscription) GetMqttSessionClientId() string`

GetMqttSessionClientId returns the MqttSessionClientId field if non-nil, zero value otherwise.

### GetMqttSessionClientIdOk

`func (o *MsgVpnMqttSessionSubscription) GetMqttSessionClientIdOk() (*string, bool)`

GetMqttSessionClientIdOk returns a tuple with the MqttSessionClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttSessionClientId

`func (o *MsgVpnMqttSessionSubscription) SetMqttSessionClientId(v string)`

SetMqttSessionClientId sets MqttSessionClientId field to given value.

### HasMqttSessionClientId

`func (o *MsgVpnMqttSessionSubscription) HasMqttSessionClientId() bool`

HasMqttSessionClientId returns a boolean if a field has been set.

### GetMqttSessionVirtualRouter

`func (o *MsgVpnMqttSessionSubscription) GetMqttSessionVirtualRouter() string`

GetMqttSessionVirtualRouter returns the MqttSessionVirtualRouter field if non-nil, zero value otherwise.

### GetMqttSessionVirtualRouterOk

`func (o *MsgVpnMqttSessionSubscription) GetMqttSessionVirtualRouterOk() (*string, bool)`

GetMqttSessionVirtualRouterOk returns a tuple with the MqttSessionVirtualRouter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttSessionVirtualRouter

`func (o *MsgVpnMqttSessionSubscription) SetMqttSessionVirtualRouter(v string)`

SetMqttSessionVirtualRouter sets MqttSessionVirtualRouter field to given value.

### HasMqttSessionVirtualRouter

`func (o *MsgVpnMqttSessionSubscription) HasMqttSessionVirtualRouter() bool`

HasMqttSessionVirtualRouter returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnMqttSessionSubscription) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnMqttSessionSubscription) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnMqttSessionSubscription) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnMqttSessionSubscription) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetSubscriptionQos

`func (o *MsgVpnMqttSessionSubscription) GetSubscriptionQos() int64`

GetSubscriptionQos returns the SubscriptionQos field if non-nil, zero value otherwise.

### GetSubscriptionQosOk

`func (o *MsgVpnMqttSessionSubscription) GetSubscriptionQosOk() (*int64, bool)`

GetSubscriptionQosOk returns a tuple with the SubscriptionQos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionQos

`func (o *MsgVpnMqttSessionSubscription) SetSubscriptionQos(v int64)`

SetSubscriptionQos sets SubscriptionQos field to given value.

### HasSubscriptionQos

`func (o *MsgVpnMqttSessionSubscription) HasSubscriptionQos() bool`

HasSubscriptionQos returns a boolean if a field has been set.

### GetSubscriptionTopic

`func (o *MsgVpnMqttSessionSubscription) GetSubscriptionTopic() string`

GetSubscriptionTopic returns the SubscriptionTopic field if non-nil, zero value otherwise.

### GetSubscriptionTopicOk

`func (o *MsgVpnMqttSessionSubscription) GetSubscriptionTopicOk() (*string, bool)`

GetSubscriptionTopicOk returns a tuple with the SubscriptionTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionTopic

`func (o *MsgVpnMqttSessionSubscription) SetSubscriptionTopic(v string)`

SetSubscriptionTopic sets SubscriptionTopic field to given value.

### HasSubscriptionTopic

`func (o *MsgVpnMqttSessionSubscription) HasSubscriptionTopic() bool`

HasSubscriptionTopic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


