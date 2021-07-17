# MsgVpnAclProfilePublishException

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclProfileName** | Pointer to **string** | The name of the ACL Profile. Deprecated since 2.14. Replaced by publishTopicExceptions. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. Deprecated since 2.14. Replaced by publishTopicExceptions. | [optional] 
**PublishExceptionTopic** | Pointer to **string** | The topic for the exception to the default action taken. May include wildcard characters. Deprecated since 2.14. Replaced by publishTopicExceptions. | [optional] 
**TopicSyntax** | Pointer to **string** | The syntax of the topic for the exception to the default action taken. The allowed values and their meaning are:  &lt;pre&gt; \&quot;smf\&quot; - Topic uses SMF syntax. \&quot;mqtt\&quot; - Topic uses MQTT syntax. &lt;/pre&gt;  Deprecated since 2.14. Replaced by publishTopicExceptions. | [optional] 

## Methods

### NewMsgVpnAclProfilePublishException

`func NewMsgVpnAclProfilePublishException() *MsgVpnAclProfilePublishException`

NewMsgVpnAclProfilePublishException instantiates a new MsgVpnAclProfilePublishException object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnAclProfilePublishExceptionWithDefaults

`func NewMsgVpnAclProfilePublishExceptionWithDefaults() *MsgVpnAclProfilePublishException`

NewMsgVpnAclProfilePublishExceptionWithDefaults instantiates a new MsgVpnAclProfilePublishException object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAclProfileName

`func (o *MsgVpnAclProfilePublishException) GetAclProfileName() string`

GetAclProfileName returns the AclProfileName field if non-nil, zero value otherwise.

### GetAclProfileNameOk

`func (o *MsgVpnAclProfilePublishException) GetAclProfileNameOk() (*string, bool)`

GetAclProfileNameOk returns a tuple with the AclProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclProfileName

`func (o *MsgVpnAclProfilePublishException) SetAclProfileName(v string)`

SetAclProfileName sets AclProfileName field to given value.

### HasAclProfileName

`func (o *MsgVpnAclProfilePublishException) HasAclProfileName() bool`

HasAclProfileName returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnAclProfilePublishException) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnAclProfilePublishException) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnAclProfilePublishException) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnAclProfilePublishException) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetPublishExceptionTopic

`func (o *MsgVpnAclProfilePublishException) GetPublishExceptionTopic() string`

GetPublishExceptionTopic returns the PublishExceptionTopic field if non-nil, zero value otherwise.

### GetPublishExceptionTopicOk

`func (o *MsgVpnAclProfilePublishException) GetPublishExceptionTopicOk() (*string, bool)`

GetPublishExceptionTopicOk returns a tuple with the PublishExceptionTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishExceptionTopic

`func (o *MsgVpnAclProfilePublishException) SetPublishExceptionTopic(v string)`

SetPublishExceptionTopic sets PublishExceptionTopic field to given value.

### HasPublishExceptionTopic

`func (o *MsgVpnAclProfilePublishException) HasPublishExceptionTopic() bool`

HasPublishExceptionTopic returns a boolean if a field has been set.

### GetTopicSyntax

`func (o *MsgVpnAclProfilePublishException) GetTopicSyntax() string`

GetTopicSyntax returns the TopicSyntax field if non-nil, zero value otherwise.

### GetTopicSyntaxOk

`func (o *MsgVpnAclProfilePublishException) GetTopicSyntaxOk() (*string, bool)`

GetTopicSyntaxOk returns a tuple with the TopicSyntax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicSyntax

`func (o *MsgVpnAclProfilePublishException) SetTopicSyntax(v string)`

SetTopicSyntax sets TopicSyntax field to given value.

### HasTopicSyntax

`func (o *MsgVpnAclProfilePublishException) HasTopicSyntax() bool`

HasTopicSyntax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


