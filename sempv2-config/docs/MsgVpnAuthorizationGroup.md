# MsgVpnAuthorizationGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclProfileName** | Pointer to **string** | The ACL Profile of the LDAP Authorization Group. The default value is &#x60;\&quot;default\&quot;&#x60;. | [optional] 
**AuthorizationGroupName** | Pointer to **string** | The name of the LDAP Authorization Group. Special care is needed if the group name contains special characters such as &#39;#&#39;, &#39;+&#39;, &#39;;&#39;, &#39;&#x3D;&#39; as the value of the group name returned from the LDAP server might prepend those characters with &#39;\\&#39;. For example a group name called &#39;test#,lab,com&#39; will be returned from the LDAP server as &#39;test\\#,lab,com&#39;. | [optional] 
**ClientProfileName** | Pointer to **string** | The Client Profile of the LDAP Authorization Group. The default value is &#x60;\&quot;default\&quot;&#x60;. | [optional] 
**Enabled** | Pointer to **bool** | Enable or disable the LDAP Authorization Group in the Message VPN. The default value is &#x60;false&#x60;. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**OrderAfterAuthorizationGroupName** | Pointer to **string** | Lower the priority to be less than this group. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. The default is not applicable. | [optional] 
**OrderBeforeAuthorizationGroupName** | Pointer to **string** | Raise the priority to be greater than this group. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. The default is not applicable. | [optional] 

## Methods

### NewMsgVpnAuthorizationGroup

`func NewMsgVpnAuthorizationGroup() *MsgVpnAuthorizationGroup`

NewMsgVpnAuthorizationGroup instantiates a new MsgVpnAuthorizationGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnAuthorizationGroupWithDefaults

`func NewMsgVpnAuthorizationGroupWithDefaults() *MsgVpnAuthorizationGroup`

NewMsgVpnAuthorizationGroupWithDefaults instantiates a new MsgVpnAuthorizationGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAclProfileName

`func (o *MsgVpnAuthorizationGroup) GetAclProfileName() string`

GetAclProfileName returns the AclProfileName field if non-nil, zero value otherwise.

### GetAclProfileNameOk

`func (o *MsgVpnAuthorizationGroup) GetAclProfileNameOk() (*string, bool)`

GetAclProfileNameOk returns a tuple with the AclProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclProfileName

`func (o *MsgVpnAuthorizationGroup) SetAclProfileName(v string)`

SetAclProfileName sets AclProfileName field to given value.

### HasAclProfileName

`func (o *MsgVpnAuthorizationGroup) HasAclProfileName() bool`

HasAclProfileName returns a boolean if a field has been set.

### GetAuthorizationGroupName

`func (o *MsgVpnAuthorizationGroup) GetAuthorizationGroupName() string`

GetAuthorizationGroupName returns the AuthorizationGroupName field if non-nil, zero value otherwise.

### GetAuthorizationGroupNameOk

`func (o *MsgVpnAuthorizationGroup) GetAuthorizationGroupNameOk() (*string, bool)`

GetAuthorizationGroupNameOk returns a tuple with the AuthorizationGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationGroupName

`func (o *MsgVpnAuthorizationGroup) SetAuthorizationGroupName(v string)`

SetAuthorizationGroupName sets AuthorizationGroupName field to given value.

### HasAuthorizationGroupName

`func (o *MsgVpnAuthorizationGroup) HasAuthorizationGroupName() bool`

HasAuthorizationGroupName returns a boolean if a field has been set.

### GetClientProfileName

`func (o *MsgVpnAuthorizationGroup) GetClientProfileName() string`

GetClientProfileName returns the ClientProfileName field if non-nil, zero value otherwise.

### GetClientProfileNameOk

`func (o *MsgVpnAuthorizationGroup) GetClientProfileNameOk() (*string, bool)`

GetClientProfileNameOk returns a tuple with the ClientProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientProfileName

`func (o *MsgVpnAuthorizationGroup) SetClientProfileName(v string)`

SetClientProfileName sets ClientProfileName field to given value.

### HasClientProfileName

`func (o *MsgVpnAuthorizationGroup) HasClientProfileName() bool`

HasClientProfileName returns a boolean if a field has been set.

### GetEnabled

`func (o *MsgVpnAuthorizationGroup) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MsgVpnAuthorizationGroup) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MsgVpnAuthorizationGroup) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MsgVpnAuthorizationGroup) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnAuthorizationGroup) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnAuthorizationGroup) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnAuthorizationGroup) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnAuthorizationGroup) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetOrderAfterAuthorizationGroupName

`func (o *MsgVpnAuthorizationGroup) GetOrderAfterAuthorizationGroupName() string`

GetOrderAfterAuthorizationGroupName returns the OrderAfterAuthorizationGroupName field if non-nil, zero value otherwise.

### GetOrderAfterAuthorizationGroupNameOk

`func (o *MsgVpnAuthorizationGroup) GetOrderAfterAuthorizationGroupNameOk() (*string, bool)`

GetOrderAfterAuthorizationGroupNameOk returns a tuple with the OrderAfterAuthorizationGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAfterAuthorizationGroupName

`func (o *MsgVpnAuthorizationGroup) SetOrderAfterAuthorizationGroupName(v string)`

SetOrderAfterAuthorizationGroupName sets OrderAfterAuthorizationGroupName field to given value.

### HasOrderAfterAuthorizationGroupName

`func (o *MsgVpnAuthorizationGroup) HasOrderAfterAuthorizationGroupName() bool`

HasOrderAfterAuthorizationGroupName returns a boolean if a field has been set.

### GetOrderBeforeAuthorizationGroupName

`func (o *MsgVpnAuthorizationGroup) GetOrderBeforeAuthorizationGroupName() string`

GetOrderBeforeAuthorizationGroupName returns the OrderBeforeAuthorizationGroupName field if non-nil, zero value otherwise.

### GetOrderBeforeAuthorizationGroupNameOk

`func (o *MsgVpnAuthorizationGroup) GetOrderBeforeAuthorizationGroupNameOk() (*string, bool)`

GetOrderBeforeAuthorizationGroupNameOk returns a tuple with the OrderBeforeAuthorizationGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBeforeAuthorizationGroupName

`func (o *MsgVpnAuthorizationGroup) SetOrderBeforeAuthorizationGroupName(v string)`

SetOrderBeforeAuthorizationGroupName sets OrderBeforeAuthorizationGroupName field to given value.

### HasOrderBeforeAuthorizationGroupName

`func (o *MsgVpnAuthorizationGroup) HasOrderBeforeAuthorizationGroupName() bool`

HasOrderBeforeAuthorizationGroupName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


