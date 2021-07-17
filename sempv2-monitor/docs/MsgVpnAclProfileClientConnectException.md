# MsgVpnAclProfileClientConnectException

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclProfileName** | Pointer to **string** | The name of the ACL Profile. | [optional] 
**ClientConnectExceptionAddress** | Pointer to **string** | The IP address/netmask of the client connect exception in CIDR form. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 

## Methods

### NewMsgVpnAclProfileClientConnectException

`func NewMsgVpnAclProfileClientConnectException() *MsgVpnAclProfileClientConnectException`

NewMsgVpnAclProfileClientConnectException instantiates a new MsgVpnAclProfileClientConnectException object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnAclProfileClientConnectExceptionWithDefaults

`func NewMsgVpnAclProfileClientConnectExceptionWithDefaults() *MsgVpnAclProfileClientConnectException`

NewMsgVpnAclProfileClientConnectExceptionWithDefaults instantiates a new MsgVpnAclProfileClientConnectException object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAclProfileName

`func (o *MsgVpnAclProfileClientConnectException) GetAclProfileName() string`

GetAclProfileName returns the AclProfileName field if non-nil, zero value otherwise.

### GetAclProfileNameOk

`func (o *MsgVpnAclProfileClientConnectException) GetAclProfileNameOk() (*string, bool)`

GetAclProfileNameOk returns a tuple with the AclProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclProfileName

`func (o *MsgVpnAclProfileClientConnectException) SetAclProfileName(v string)`

SetAclProfileName sets AclProfileName field to given value.

### HasAclProfileName

`func (o *MsgVpnAclProfileClientConnectException) HasAclProfileName() bool`

HasAclProfileName returns a boolean if a field has been set.

### GetClientConnectExceptionAddress

`func (o *MsgVpnAclProfileClientConnectException) GetClientConnectExceptionAddress() string`

GetClientConnectExceptionAddress returns the ClientConnectExceptionAddress field if non-nil, zero value otherwise.

### GetClientConnectExceptionAddressOk

`func (o *MsgVpnAclProfileClientConnectException) GetClientConnectExceptionAddressOk() (*string, bool)`

GetClientConnectExceptionAddressOk returns a tuple with the ClientConnectExceptionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientConnectExceptionAddress

`func (o *MsgVpnAclProfileClientConnectException) SetClientConnectExceptionAddress(v string)`

SetClientConnectExceptionAddress sets ClientConnectExceptionAddress field to given value.

### HasClientConnectExceptionAddress

`func (o *MsgVpnAclProfileClientConnectException) HasClientConnectExceptionAddress() bool`

HasClientConnectExceptionAddress returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnAclProfileClientConnectException) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnAclProfileClientConnectException) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnAclProfileClientConnectException) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnAclProfileClientConnectException) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


