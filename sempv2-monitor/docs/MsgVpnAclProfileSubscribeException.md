# MsgVpnAclProfileSubscribeException

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclProfileName** | Pointer to **string** | The name of the ACL Profile. Deprecated since 2.14. Replaced by subscribeTopicExceptions. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. Deprecated since 2.14. Replaced by subscribeTopicExceptions. | [optional] 
**SubscribeExceptionTopic** | Pointer to **string** | The topic for the exception to the default action taken. May include wildcard characters. Deprecated since 2.14. Replaced by subscribeTopicExceptions. | [optional] 
**TopicSyntax** | Pointer to **string** | The syntax of the topic for the exception to the default action taken. The allowed values and their meaning are:  &lt;pre&gt; \&quot;smf\&quot; - Topic uses SMF syntax. \&quot;mqtt\&quot; - Topic uses MQTT syntax. &lt;/pre&gt;  Deprecated since 2.14. Replaced by subscribeTopicExceptions. | [optional] 

## Methods

### NewMsgVpnAclProfileSubscribeException

`func NewMsgVpnAclProfileSubscribeException() *MsgVpnAclProfileSubscribeException`

NewMsgVpnAclProfileSubscribeException instantiates a new MsgVpnAclProfileSubscribeException object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnAclProfileSubscribeExceptionWithDefaults

`func NewMsgVpnAclProfileSubscribeExceptionWithDefaults() *MsgVpnAclProfileSubscribeException`

NewMsgVpnAclProfileSubscribeExceptionWithDefaults instantiates a new MsgVpnAclProfileSubscribeException object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAclProfileName

`func (o *MsgVpnAclProfileSubscribeException) GetAclProfileName() string`

GetAclProfileName returns the AclProfileName field if non-nil, zero value otherwise.

### GetAclProfileNameOk

`func (o *MsgVpnAclProfileSubscribeException) GetAclProfileNameOk() (*string, bool)`

GetAclProfileNameOk returns a tuple with the AclProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclProfileName

`func (o *MsgVpnAclProfileSubscribeException) SetAclProfileName(v string)`

SetAclProfileName sets AclProfileName field to given value.

### HasAclProfileName

`func (o *MsgVpnAclProfileSubscribeException) HasAclProfileName() bool`

HasAclProfileName returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnAclProfileSubscribeException) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnAclProfileSubscribeException) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnAclProfileSubscribeException) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnAclProfileSubscribeException) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetSubscribeExceptionTopic

`func (o *MsgVpnAclProfileSubscribeException) GetSubscribeExceptionTopic() string`

GetSubscribeExceptionTopic returns the SubscribeExceptionTopic field if non-nil, zero value otherwise.

### GetSubscribeExceptionTopicOk

`func (o *MsgVpnAclProfileSubscribeException) GetSubscribeExceptionTopicOk() (*string, bool)`

GetSubscribeExceptionTopicOk returns a tuple with the SubscribeExceptionTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeExceptionTopic

`func (o *MsgVpnAclProfileSubscribeException) SetSubscribeExceptionTopic(v string)`

SetSubscribeExceptionTopic sets SubscribeExceptionTopic field to given value.

### HasSubscribeExceptionTopic

`func (o *MsgVpnAclProfileSubscribeException) HasSubscribeExceptionTopic() bool`

HasSubscribeExceptionTopic returns a boolean if a field has been set.

### GetTopicSyntax

`func (o *MsgVpnAclProfileSubscribeException) GetTopicSyntax() string`

GetTopicSyntax returns the TopicSyntax field if non-nil, zero value otherwise.

### GetTopicSyntaxOk

`func (o *MsgVpnAclProfileSubscribeException) GetTopicSyntaxOk() (*string, bool)`

GetTopicSyntaxOk returns a tuple with the TopicSyntax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicSyntax

`func (o *MsgVpnAclProfileSubscribeException) SetTopicSyntax(v string)`

SetTopicSyntax sets TopicSyntax field to given value.

### HasTopicSyntax

`func (o *MsgVpnAclProfileSubscribeException) HasTopicSyntax() bool`

HasTopicSyntax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


