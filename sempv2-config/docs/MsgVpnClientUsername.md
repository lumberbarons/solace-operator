# MsgVpnClientUsername

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclProfileName** | Pointer to **string** | The ACL Profile of the Client Username. The default value is &#x60;\&quot;default\&quot;&#x60;. | [optional] 
**ClientProfileName** | Pointer to **string** | The Client Profile of the Client Username. The default value is &#x60;\&quot;default\&quot;&#x60;. | [optional] 
**ClientUsername** | Pointer to **string** | The name of the Client Username. | [optional] 
**Enabled** | Pointer to **bool** | Enable or disable the Client Username. When disabled, all clients currently connected as the Client Username are disconnected. The default value is &#x60;false&#x60;. | [optional] 
**GuaranteedEndpointPermissionOverrideEnabled** | Pointer to **bool** | Enable or disable guaranteed endpoint permission override for the Client Username. When enabled all guaranteed endpoints may be accessed, modified or deleted with the same permission as the owner. The default value is &#x60;false&#x60;. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**Password** | Pointer to **string** | The password for the Client Username. This attribute is absent from a GET and not updated when absent in a PUT, subject to the exceptions in note 4. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] 
**SubscriptionManagerEnabled** | Pointer to **bool** | Enable or disable the subscription management capability of the Client Username. This is the ability to manage subscriptions on behalf of other Client Usernames. The default value is &#x60;false&#x60;. | [optional] 

## Methods

### NewMsgVpnClientUsername

`func NewMsgVpnClientUsername() *MsgVpnClientUsername`

NewMsgVpnClientUsername instantiates a new MsgVpnClientUsername object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnClientUsernameWithDefaults

`func NewMsgVpnClientUsernameWithDefaults() *MsgVpnClientUsername`

NewMsgVpnClientUsernameWithDefaults instantiates a new MsgVpnClientUsername object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAclProfileName

`func (o *MsgVpnClientUsername) GetAclProfileName() string`

GetAclProfileName returns the AclProfileName field if non-nil, zero value otherwise.

### GetAclProfileNameOk

`func (o *MsgVpnClientUsername) GetAclProfileNameOk() (*string, bool)`

GetAclProfileNameOk returns a tuple with the AclProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclProfileName

`func (o *MsgVpnClientUsername) SetAclProfileName(v string)`

SetAclProfileName sets AclProfileName field to given value.

### HasAclProfileName

`func (o *MsgVpnClientUsername) HasAclProfileName() bool`

HasAclProfileName returns a boolean if a field has been set.

### GetClientProfileName

`func (o *MsgVpnClientUsername) GetClientProfileName() string`

GetClientProfileName returns the ClientProfileName field if non-nil, zero value otherwise.

### GetClientProfileNameOk

`func (o *MsgVpnClientUsername) GetClientProfileNameOk() (*string, bool)`

GetClientProfileNameOk returns a tuple with the ClientProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientProfileName

`func (o *MsgVpnClientUsername) SetClientProfileName(v string)`

SetClientProfileName sets ClientProfileName field to given value.

### HasClientProfileName

`func (o *MsgVpnClientUsername) HasClientProfileName() bool`

HasClientProfileName returns a boolean if a field has been set.

### GetClientUsername

`func (o *MsgVpnClientUsername) GetClientUsername() string`

GetClientUsername returns the ClientUsername field if non-nil, zero value otherwise.

### GetClientUsernameOk

`func (o *MsgVpnClientUsername) GetClientUsernameOk() (*string, bool)`

GetClientUsernameOk returns a tuple with the ClientUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUsername

`func (o *MsgVpnClientUsername) SetClientUsername(v string)`

SetClientUsername sets ClientUsername field to given value.

### HasClientUsername

`func (o *MsgVpnClientUsername) HasClientUsername() bool`

HasClientUsername returns a boolean if a field has been set.

### GetEnabled

`func (o *MsgVpnClientUsername) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MsgVpnClientUsername) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MsgVpnClientUsername) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MsgVpnClientUsername) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGuaranteedEndpointPermissionOverrideEnabled

`func (o *MsgVpnClientUsername) GetGuaranteedEndpointPermissionOverrideEnabled() bool`

GetGuaranteedEndpointPermissionOverrideEnabled returns the GuaranteedEndpointPermissionOverrideEnabled field if non-nil, zero value otherwise.

### GetGuaranteedEndpointPermissionOverrideEnabledOk

`func (o *MsgVpnClientUsername) GetGuaranteedEndpointPermissionOverrideEnabledOk() (*bool, bool)`

GetGuaranteedEndpointPermissionOverrideEnabledOk returns a tuple with the GuaranteedEndpointPermissionOverrideEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteedEndpointPermissionOverrideEnabled

`func (o *MsgVpnClientUsername) SetGuaranteedEndpointPermissionOverrideEnabled(v bool)`

SetGuaranteedEndpointPermissionOverrideEnabled sets GuaranteedEndpointPermissionOverrideEnabled field to given value.

### HasGuaranteedEndpointPermissionOverrideEnabled

`func (o *MsgVpnClientUsername) HasGuaranteedEndpointPermissionOverrideEnabled() bool`

HasGuaranteedEndpointPermissionOverrideEnabled returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnClientUsername) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnClientUsername) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnClientUsername) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnClientUsername) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetPassword

`func (o *MsgVpnClientUsername) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MsgVpnClientUsername) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MsgVpnClientUsername) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *MsgVpnClientUsername) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSubscriptionManagerEnabled

`func (o *MsgVpnClientUsername) GetSubscriptionManagerEnabled() bool`

GetSubscriptionManagerEnabled returns the SubscriptionManagerEnabled field if non-nil, zero value otherwise.

### GetSubscriptionManagerEnabledOk

`func (o *MsgVpnClientUsername) GetSubscriptionManagerEnabledOk() (*bool, bool)`

GetSubscriptionManagerEnabledOk returns a tuple with the SubscriptionManagerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionManagerEnabled

`func (o *MsgVpnClientUsername) SetSubscriptionManagerEnabled(v bool)`

SetSubscriptionManagerEnabled sets SubscriptionManagerEnabled field to given value.

### HasSubscriptionManagerEnabled

`func (o *MsgVpnClientUsername) HasSubscriptionManagerEnabled() bool`

HasSubscriptionManagerEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


