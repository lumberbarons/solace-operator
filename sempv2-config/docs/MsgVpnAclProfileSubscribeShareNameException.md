# MsgVpnAclProfileSubscribeShareNameException

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclProfileName** | Pointer to **string** | The name of the ACL Profile. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**SubscribeShareNameException** | Pointer to **string** | The subscribe share name exception to the default action taken. May include wildcard characters. | [optional] 
**SubscribeShareNameExceptionSyntax** | Pointer to **string** | The syntax of the subscribe share name for the exception to the default action taken. The allowed values and their meaning are:  &lt;pre&gt; \&quot;smf\&quot; - Topic uses SMF syntax. \&quot;mqtt\&quot; - Topic uses MQTT syntax. &lt;/pre&gt;  | [optional] 

## Methods

### NewMsgVpnAclProfileSubscribeShareNameException

`func NewMsgVpnAclProfileSubscribeShareNameException() *MsgVpnAclProfileSubscribeShareNameException`

NewMsgVpnAclProfileSubscribeShareNameException instantiates a new MsgVpnAclProfileSubscribeShareNameException object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnAclProfileSubscribeShareNameExceptionWithDefaults

`func NewMsgVpnAclProfileSubscribeShareNameExceptionWithDefaults() *MsgVpnAclProfileSubscribeShareNameException`

NewMsgVpnAclProfileSubscribeShareNameExceptionWithDefaults instantiates a new MsgVpnAclProfileSubscribeShareNameException object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAclProfileName

`func (o *MsgVpnAclProfileSubscribeShareNameException) GetAclProfileName() string`

GetAclProfileName returns the AclProfileName field if non-nil, zero value otherwise.

### GetAclProfileNameOk

`func (o *MsgVpnAclProfileSubscribeShareNameException) GetAclProfileNameOk() (*string, bool)`

GetAclProfileNameOk returns a tuple with the AclProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclProfileName

`func (o *MsgVpnAclProfileSubscribeShareNameException) SetAclProfileName(v string)`

SetAclProfileName sets AclProfileName field to given value.

### HasAclProfileName

`func (o *MsgVpnAclProfileSubscribeShareNameException) HasAclProfileName() bool`

HasAclProfileName returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnAclProfileSubscribeShareNameException) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnAclProfileSubscribeShareNameException) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnAclProfileSubscribeShareNameException) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnAclProfileSubscribeShareNameException) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetSubscribeShareNameException

`func (o *MsgVpnAclProfileSubscribeShareNameException) GetSubscribeShareNameException() string`

GetSubscribeShareNameException returns the SubscribeShareNameException field if non-nil, zero value otherwise.

### GetSubscribeShareNameExceptionOk

`func (o *MsgVpnAclProfileSubscribeShareNameException) GetSubscribeShareNameExceptionOk() (*string, bool)`

GetSubscribeShareNameExceptionOk returns a tuple with the SubscribeShareNameException field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeShareNameException

`func (o *MsgVpnAclProfileSubscribeShareNameException) SetSubscribeShareNameException(v string)`

SetSubscribeShareNameException sets SubscribeShareNameException field to given value.

### HasSubscribeShareNameException

`func (o *MsgVpnAclProfileSubscribeShareNameException) HasSubscribeShareNameException() bool`

HasSubscribeShareNameException returns a boolean if a field has been set.

### GetSubscribeShareNameExceptionSyntax

`func (o *MsgVpnAclProfileSubscribeShareNameException) GetSubscribeShareNameExceptionSyntax() string`

GetSubscribeShareNameExceptionSyntax returns the SubscribeShareNameExceptionSyntax field if non-nil, zero value otherwise.

### GetSubscribeShareNameExceptionSyntaxOk

`func (o *MsgVpnAclProfileSubscribeShareNameException) GetSubscribeShareNameExceptionSyntaxOk() (*string, bool)`

GetSubscribeShareNameExceptionSyntaxOk returns a tuple with the SubscribeShareNameExceptionSyntax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribeShareNameExceptionSyntax

`func (o *MsgVpnAclProfileSubscribeShareNameException) SetSubscribeShareNameExceptionSyntax(v string)`

SetSubscribeShareNameExceptionSyntax sets SubscribeShareNameExceptionSyntax field to given value.

### HasSubscribeShareNameExceptionSyntax

`func (o *MsgVpnAclProfileSubscribeShareNameException) HasSubscribeShareNameExceptionSyntax() bool`

HasSubscribeShareNameExceptionSyntax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


