# MsgVpnAclProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclProfileName** | Pointer to **string** | The name of the ACL Profile. | [optional] 
**ClientConnectDefaultAction** | Pointer to **string** | The default action to take when a client using the ACL Profile connects to the Message VPN. The allowed values and their meaning are:  &lt;pre&gt; \&quot;allow\&quot; - Allow client connection unless an exception is found for it. \&quot;disallow\&quot; - Disallow client connection unless an exception is found for it. &lt;/pre&gt;  | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**PublishTopicDefaultAction** | Pointer to **string** | The default action to take when a client using the ACL Profile publishes to a topic in the Message VPN. The allowed values and their meaning are:  &lt;pre&gt; \&quot;allow\&quot; - Allow topic unless an exception is found for it. \&quot;disallow\&quot; - Disallow topic unless an exception is found for it. &lt;/pre&gt;  | [optional] 
**SubscribeShareNameDefaultAction** | Pointer to **string** | The default action to take when a client using the ACL Profile subscribes to a share-name subscription in the Message VPN. The allowed values and their meaning are:  &lt;pre&gt; \&quot;allow\&quot; - Allow topic unless an exception is found for it. \&quot;disallow\&quot; - Disallow topic unless an exception is found for it. &lt;/pre&gt;  Available since 2.14. | [optional] 
**SubscribeTopicDefaultAction** | Pointer to **string** | The default action to take when a client using the ACL Profile subscribes to a topic in the Message VPN. The allowed values and their meaning are:  &lt;pre&gt; \&quot;allow\&quot; - Allow topic unless an exception is found for it. \&quot;disallow\&quot; - Disallow topic unless an exception is found for it. &lt;/pre&gt;  | [optional] 

## Methods

### NewMsgVpnAclProfile

`func NewMsgVpnAclProfile() *MsgVpnAclProfile`

NewMsgVpnAclProfile instantiates a new MsgVpnAclProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnAclProfileWithDefaults

`func NewMsgVpnAclProfileWithDefaults() *MsgVpnAclProfile`

NewMsgVpnAclProfileWithDefaults instantiates a new MsgVpnAclProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAclProfileName

`func (o *MsgVpnAclProfile) GetAclProfileName() string`

GetAclProfileName returns the AclProfileName field if non-nil, zero value otherwise.

### GetAclProfileNameOk

`func (o *MsgVpnAclProfile) GetAclProfileNameOk() (*string, bool)`

GetAclProfileNameOk returns a tuple with the AclProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclProfileName

`func (o *MsgVpnAclProfile) SetAclProfileName(v string)`

SetAclProfileName sets AclProfileName field to given value.

### HasAclProfileName

`func (o *MsgVpnAclProfile) HasAclProfileName() bool`

HasAclProfileName returns a boolean if a field has been set.

### GetClientConnectDefaultAction

`func (o *MsgVpnAclProfile) GetClientConnectDefaultAction() string`

GetClientConnectDefaultAction returns the ClientConnectDefaultAction field if non-nil, zero value otherwise.

### GetClientConnectDefaultActionOk

`func (o *MsgVpnAclProfile) GetClientConnectDefaultActionOk() (*string, bool)`

GetClientConnectDefaultActionOk returns a tuple with the ClientConnectDefaultAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientConnectDefaultAction

`func (o *MsgVpnAclProfile) SetClientConnectDefaultAction(v string)`

SetClientConnectDefaultAction sets ClientConnectDefaultAction field to given value.

### HasClientConnectDefaultAction

`func (o *MsgVpnAclProfile) HasClientConnectDefaultAction() bool`

HasClientConnectDefaultAction returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnAclProfile) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnAclProfile) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnAclProfile) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnAclProfile) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetPublishTopicDefaultAction

`func (o *MsgVpnAclProfile) GetPublishTopicDefaultAction() string`

GetPublishTopicDefaultAction returns the PublishTopicDefaultAction field if non-nil, zero value otherwise.

### GetPublishTopicDefaultActionOk

`func (o *MsgVpnAclProfile) GetPublishTopicDefaultActionOk() (*string, bool)`

GetPublishTopicDefaultActionOk returns a tuple with the PublishTopicDefaultAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishTopicDefaultAction

`func (o *MsgVpnAclProfile) SetPublishTopicDefaultAction(v string)`

SetPublishTopicDefaultAction sets PublishTopicDefaultAction field to given value.

### HasPublishTopicDefaultAction

`func (o *MsgVpnAclProfile) HasPublishTopicDefaultAction() bool`

HasPublishTopicDefaultAction returns a boolean if a field has been set.

### GetSubscribeShareNameDefaultAction

`func (o *MsgVpnAclProfile) GetSubscribeShareNameDefaultAction() string`

GetSubscribeShareNameDefaultAction returns the SubscribeShareNameDefaultAction field if non-nil, zero value otherwise.

### GetSubscribeShareNameDefaultActionOk

`func (o *MsgVpnAclProfile) GetSubscribeShareNameDefaultActionOk() (*string, bool)`

GetSubscribeShareNameDefaultActionOk returns a tuple with the SubscribeShareNameDefaultAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeShareNameDefaultAction

`func (o *MsgVpnAclProfile) SetSubscribeShareNameDefaultAction(v string)`

SetSubscribeShareNameDefaultAction sets SubscribeShareNameDefaultAction field to given value.

### HasSubscribeShareNameDefaultAction

`func (o *MsgVpnAclProfile) HasSubscribeShareNameDefaultAction() bool`

HasSubscribeShareNameDefaultAction returns a boolean if a field has been set.

### GetSubscribeTopicDefaultAction

`func (o *MsgVpnAclProfile) GetSubscribeTopicDefaultAction() string`

GetSubscribeTopicDefaultAction returns the SubscribeTopicDefaultAction field if non-nil, zero value otherwise.

### GetSubscribeTopicDefaultActionOk

`func (o *MsgVpnAclProfile) GetSubscribeTopicDefaultActionOk() (*string, bool)`

GetSubscribeTopicDefaultActionOk returns a tuple with the SubscribeTopicDefaultAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeTopicDefaultAction

`func (o *MsgVpnAclProfile) SetSubscribeTopicDefaultAction(v string)`

SetSubscribeTopicDefaultAction sets SubscribeTopicDefaultAction field to given value.

### HasSubscribeTopicDefaultAction

`func (o *MsgVpnAclProfile) HasSubscribeTopicDefaultAction() bool`

HasSubscribeTopicDefaultAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


