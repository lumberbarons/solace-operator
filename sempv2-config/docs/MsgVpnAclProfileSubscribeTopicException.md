# MsgVpnAclProfileSubscribeTopicException

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclProfileName** | Pointer to **string** | The name of the ACL Profile. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**SubscribeTopicException** | Pointer to **string** | The topic for the exception to the default action taken. May include wildcard characters. | [optional] 
**SubscribeTopicExceptionSyntax** | Pointer to **string** | The syntax of the topic for the exception to the default action taken. The allowed values and their meaning are:  &lt;pre&gt; \&quot;smf\&quot; - Topic uses SMF syntax. \&quot;mqtt\&quot; - Topic uses MQTT syntax. &lt;/pre&gt;  | [optional] 

## Methods

### NewMsgVpnAclProfileSubscribeTopicException

`func NewMsgVpnAclProfileSubscribeTopicException() *MsgVpnAclProfileSubscribeTopicException`

NewMsgVpnAclProfileSubscribeTopicException instantiates a new MsgVpnAclProfileSubscribeTopicException object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnAclProfileSubscribeTopicExceptionWithDefaults

`func NewMsgVpnAclProfileSubscribeTopicExceptionWithDefaults() *MsgVpnAclProfileSubscribeTopicException`

NewMsgVpnAclProfileSubscribeTopicExceptionWithDefaults instantiates a new MsgVpnAclProfileSubscribeTopicException object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAclProfileName

`func (o *MsgVpnAclProfileSubscribeTopicException) GetAclProfileName() string`

GetAclProfileName returns the AclProfileName field if non-nil, zero value otherwise.

### GetAclProfileNameOk

`func (o *MsgVpnAclProfileSubscribeTopicException) GetAclProfileNameOk() (*string, bool)`

GetAclProfileNameOk returns a tuple with the AclProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclProfileName

`func (o *MsgVpnAclProfileSubscribeTopicException) SetAclProfileName(v string)`

SetAclProfileName sets AclProfileName field to given value.

### HasAclProfileName

`func (o *MsgVpnAclProfileSubscribeTopicException) HasAclProfileName() bool`

HasAclProfileName returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnAclProfileSubscribeTopicException) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnAclProfileSubscribeTopicException) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnAclProfileSubscribeTopicException) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnAclProfileSubscribeTopicException) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetSubscribeTopicException

`func (o *MsgVpnAclProfileSubscribeTopicException) GetSubscribeTopicException() string`

GetSubscribeTopicException returns the SubscribeTopicException field if non-nil, zero value otherwise.

### GetSubscribeTopicExceptionOk

`func (o *MsgVpnAclProfileSubscribeTopicException) GetSubscribeTopicExceptionOk() (*string, bool)`

GetSubscribeTopicExceptionOk returns a tuple with the SubscribeTopicException field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeTopicException

`func (o *MsgVpnAclProfileSubscribeTopicException) SetSubscribeTopicException(v string)`

SetSubscribeTopicException sets SubscribeTopicException field to given value.

### HasSubscribeTopicException

`func (o *MsgVpnAclProfileSubscribeTopicException) HasSubscribeTopicException() bool`

HasSubscribeTopicException returns a boolean if a field has been set.

### GetSubscribeTopicExceptionSyntax

`func (o *MsgVpnAclProfileSubscribeTopicException) GetSubscribeTopicExceptionSyntax() string`

GetSubscribeTopicExceptionSyntax returns the SubscribeTopicExceptionSyntax field if non-nil, zero value otherwise.

### GetSubscribeTopicExceptionSyntaxOk

`func (o *MsgVpnAclProfileSubscribeTopicException) GetSubscribeTopicExceptionSyntaxOk() (*string, bool)`

GetSubscribeTopicExceptionSyntaxOk returns a tuple with the SubscribeTopicExceptionSyntax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeTopicExceptionSyntax

`func (o *MsgVpnAclProfileSubscribeTopicException) SetSubscribeTopicExceptionSyntax(v string)`

SetSubscribeTopicExceptionSyntax sets SubscribeTopicExceptionSyntax field to given value.

### HasSubscribeTopicExceptionSyntax

`func (o *MsgVpnAclProfileSubscribeTopicException) HasSubscribeTopicExceptionSyntax() bool`

HasSubscribeTopicExceptionSyntax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


