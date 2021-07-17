# MsgVpnAclProfilePublishTopicException

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclProfileName** | Pointer to **string** | The name of the ACL Profile. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**PublishTopicException** | Pointer to **string** | The topic for the exception to the default action taken. May include wildcard characters. | [optional] 
**PublishTopicExceptionSyntax** | Pointer to **string** | The syntax of the topic for the exception to the default action taken. The allowed values and their meaning are:  &lt;pre&gt; \&quot;smf\&quot; - Topic uses SMF syntax. \&quot;mqtt\&quot; - Topic uses MQTT syntax. &lt;/pre&gt;  | [optional] 

## Methods

### NewMsgVpnAclProfilePublishTopicException

`func NewMsgVpnAclProfilePublishTopicException() *MsgVpnAclProfilePublishTopicException`

NewMsgVpnAclProfilePublishTopicException instantiates a new MsgVpnAclProfilePublishTopicException object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnAclProfilePublishTopicExceptionWithDefaults

`func NewMsgVpnAclProfilePublishTopicExceptionWithDefaults() *MsgVpnAclProfilePublishTopicException`

NewMsgVpnAclProfilePublishTopicExceptionWithDefaults instantiates a new MsgVpnAclProfilePublishTopicException object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAclProfileName

`func (o *MsgVpnAclProfilePublishTopicException) GetAclProfileName() string`

GetAclProfileName returns the AclProfileName field if non-nil, zero value otherwise.

### GetAclProfileNameOk

`func (o *MsgVpnAclProfilePublishTopicException) GetAclProfileNameOk() (*string, bool)`

GetAclProfileNameOk returns a tuple with the AclProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclProfileName

`func (o *MsgVpnAclProfilePublishTopicException) SetAclProfileName(v string)`

SetAclProfileName sets AclProfileName field to given value.

### HasAclProfileName

`func (o *MsgVpnAclProfilePublishTopicException) HasAclProfileName() bool`

HasAclProfileName returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnAclProfilePublishTopicException) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnAclProfilePublishTopicException) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnAclProfilePublishTopicException) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnAclProfilePublishTopicException) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetPublishTopicException

`func (o *MsgVpnAclProfilePublishTopicException) GetPublishTopicException() string`

GetPublishTopicException returns the PublishTopicException field if non-nil, zero value otherwise.

### GetPublishTopicExceptionOk

`func (o *MsgVpnAclProfilePublishTopicException) GetPublishTopicExceptionOk() (*string, bool)`

GetPublishTopicExceptionOk returns a tuple with the PublishTopicException field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishTopicException

`func (o *MsgVpnAclProfilePublishTopicException) SetPublishTopicException(v string)`

SetPublishTopicException sets PublishTopicException field to given value.

### HasPublishTopicException

`func (o *MsgVpnAclProfilePublishTopicException) HasPublishTopicException() bool`

HasPublishTopicException returns a boolean if a field has been set.

### GetPublishTopicExceptionSyntax

`func (o *MsgVpnAclProfilePublishTopicException) GetPublishTopicExceptionSyntax() string`

GetPublishTopicExceptionSyntax returns the PublishTopicExceptionSyntax field if non-nil, zero value otherwise.

### GetPublishTopicExceptionSyntaxOk

`func (o *MsgVpnAclProfilePublishTopicException) GetPublishTopicExceptionSyntaxOk() (*string, bool)`

GetPublishTopicExceptionSyntaxOk returns a tuple with the PublishTopicExceptionSyntax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishTopicExceptionSyntax

`func (o *MsgVpnAclProfilePublishTopicException) SetPublishTopicExceptionSyntax(v string)`

SetPublishTopicExceptionSyntax sets PublishTopicExceptionSyntax field to given value.

### HasPublishTopicExceptionSyntax

`func (o *MsgVpnAclProfilePublishTopicException) HasPublishTopicExceptionSyntax() bool`

HasPublishTopicExceptionSyntax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


