# MsgVpnRestDeliveryPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientProfileName** | Pointer to **string** | The Client Profile of the REST Delivery Point. It must exist in the local Message VPN. Its TCP parameters are used for all REST Consumers in this RDP. Its queue properties are used by the RDP client. The Client Profile is used inside the auto-generated Client Username for this RDP. The default value is &#x60;\&quot;default\&quot;&#x60;. | [optional] 
**Enabled** | Pointer to **bool** | Enable or disable the REST Delivery Point. When disabled, no connections are initiated or messages delivered to any of the contained REST Consumers. The default value is &#x60;false&#x60;. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**RestDeliveryPointName** | Pointer to **string** | The name of the REST Delivery Point. | [optional] 
**Service** | Pointer to **string** | The name of the service that this REST Delivery Point connects to. Internally the broker does not use this value; it is informational only. The default value is &#x60;\&quot;\&quot;&#x60;. Available since 2.19. | [optional] 
**Vendor** | Pointer to **string** | The name of the vendor that this REST Delivery Point connects to. Internally the broker does not use this value; it is informational only. The default value is &#x60;\&quot;\&quot;&#x60;. Available since 2.19. | [optional] 

## Methods

### NewMsgVpnRestDeliveryPoint

`func NewMsgVpnRestDeliveryPoint() *MsgVpnRestDeliveryPoint`

NewMsgVpnRestDeliveryPoint instantiates a new MsgVpnRestDeliveryPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnRestDeliveryPointWithDefaults

`func NewMsgVpnRestDeliveryPointWithDefaults() *MsgVpnRestDeliveryPoint`

NewMsgVpnRestDeliveryPointWithDefaults instantiates a new MsgVpnRestDeliveryPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientProfileName

`func (o *MsgVpnRestDeliveryPoint) GetClientProfileName() string`

GetClientProfileName returns the ClientProfileName field if non-nil, zero value otherwise.

### GetClientProfileNameOk

`func (o *MsgVpnRestDeliveryPoint) GetClientProfileNameOk() (*string, bool)`

GetClientProfileNameOk returns a tuple with the ClientProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientProfileName

`func (o *MsgVpnRestDeliveryPoint) SetClientProfileName(v string)`

SetClientProfileName sets ClientProfileName field to given value.

### HasClientProfileName

`func (o *MsgVpnRestDeliveryPoint) HasClientProfileName() bool`

HasClientProfileName returns a boolean if a field has been set.

### GetEnabled

`func (o *MsgVpnRestDeliveryPoint) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MsgVpnRestDeliveryPoint) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MsgVpnRestDeliveryPoint) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MsgVpnRestDeliveryPoint) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnRestDeliveryPoint) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnRestDeliveryPoint) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnRestDeliveryPoint) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnRestDeliveryPoint) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetRestDeliveryPointName

`func (o *MsgVpnRestDeliveryPoint) GetRestDeliveryPointName() string`

GetRestDeliveryPointName returns the RestDeliveryPointName field if non-nil, zero value otherwise.

### GetRestDeliveryPointNameOk

`func (o *MsgVpnRestDeliveryPoint) GetRestDeliveryPointNameOk() (*string, bool)`

GetRestDeliveryPointNameOk returns a tuple with the RestDeliveryPointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestDeliveryPointName

`func (o *MsgVpnRestDeliveryPoint) SetRestDeliveryPointName(v string)`

SetRestDeliveryPointName sets RestDeliveryPointName field to given value.

### HasRestDeliveryPointName

`func (o *MsgVpnRestDeliveryPoint) HasRestDeliveryPointName() bool`

HasRestDeliveryPointName returns a boolean if a field has been set.

### GetService

`func (o *MsgVpnRestDeliveryPoint) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *MsgVpnRestDeliveryPoint) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *MsgVpnRestDeliveryPoint) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *MsgVpnRestDeliveryPoint) HasService() bool`

HasService returns a boolean if a field has been set.

### GetVendor

`func (o *MsgVpnRestDeliveryPoint) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *MsgVpnRestDeliveryPoint) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *MsgVpnRestDeliveryPoint) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *MsgVpnRestDeliveryPoint) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


