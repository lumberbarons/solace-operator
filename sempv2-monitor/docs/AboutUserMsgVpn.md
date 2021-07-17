# AboutUserMsgVpn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessLevel** | Pointer to **string** | The Message VPN access level of the User. The allowed values and their meaning are:  &lt;pre&gt; \&quot;none\&quot; - No access. \&quot;read-only\&quot; - Read only access. \&quot;read-write\&quot; - Read and write access. &lt;/pre&gt;  | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 

## Methods

### NewAboutUserMsgVpn

`func NewAboutUserMsgVpn() *AboutUserMsgVpn`

NewAboutUserMsgVpn instantiates a new AboutUserMsgVpn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAboutUserMsgVpnWithDefaults

`func NewAboutUserMsgVpnWithDefaults() *AboutUserMsgVpn`

NewAboutUserMsgVpnWithDefaults instantiates a new AboutUserMsgVpn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessLevel

`func (o *AboutUserMsgVpn) GetAccessLevel() string`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *AboutUserMsgVpn) GetAccessLevelOk() (*string, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *AboutUserMsgVpn) SetAccessLevel(v string)`

SetAccessLevel sets AccessLevel field to given value.

### HasAccessLevel

`func (o *AboutUserMsgVpn) HasAccessLevel() bool`

HasAccessLevel returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *AboutUserMsgVpn) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *AboutUserMsgVpn) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *AboutUserMsgVpn) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *AboutUserMsgVpn) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


