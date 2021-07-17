# VirtualHostname

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enable or disable Virtual Hostname to Message VPN mapping. | [optional] 
**MsgVpnName** | Pointer to **string** | The message VPN to which this virtual hostname is mapped. | [optional] 
**VirtualHostname** | Pointer to **string** | The virtual hostname. | [optional] 

## Methods

### NewVirtualHostname

`func NewVirtualHostname() *VirtualHostname`

NewVirtualHostname instantiates a new VirtualHostname object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualHostnameWithDefaults

`func NewVirtualHostnameWithDefaults() *VirtualHostname`

NewVirtualHostnameWithDefaults instantiates a new VirtualHostname object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *VirtualHostname) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VirtualHostname) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VirtualHostname) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *VirtualHostname) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *VirtualHostname) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *VirtualHostname) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *VirtualHostname) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *VirtualHostname) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetVirtualHostname

`func (o *VirtualHostname) GetVirtualHostname() string`

GetVirtualHostname returns the VirtualHostname field if non-nil, zero value otherwise.

### GetVirtualHostnameOk

`func (o *VirtualHostname) GetVirtualHostnameOk() (*string, bool)`

GetVirtualHostnameOk returns a tuple with the VirtualHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualHostname

`func (o *VirtualHostname) SetVirtualHostname(v string)`

SetVirtualHostname sets VirtualHostname field to given value.

### HasVirtualHostname

`func (o *VirtualHostname) HasVirtualHostname() bool`

HasVirtualHostname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


