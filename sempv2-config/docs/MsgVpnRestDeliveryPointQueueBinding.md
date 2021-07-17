# MsgVpnRestDeliveryPointQueueBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayReplaceTargetAuthorityEnabled** | Pointer to **bool** | Enable or disable whether the authority for the request-target is replaced with that configured for the REST Consumer remote. When enabled, the broker sends HTTP requests in absolute-form, with the request-target&#39;s authority taken from the REST Consumer&#39;s remote host and port configuration. When disabled, the broker sends HTTP requests whose request-target matches that of the original request message, including whether to use absolute-form or origin-form. This configuration is applicable only when the Message VPN is in REST gateway mode. The default value is &#x60;false&#x60;. Available since 2.6. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**PostRequestTarget** | Pointer to **string** | The request-target string to use when sending requests. It identifies the target resource on the far-end REST Consumer upon which to apply the request. There are generally two common forms for the request-target. The origin-form is most often used in practice and contains the path and query components of the target URI. If the path component is empty then the client must generally send a \&quot;/\&quot; as the path. When making a request to a proxy, most often the absolute-form is required. This configuration is only applicable when the Message VPN is in REST messaging mode. The default value is &#x60;\&quot;\&quot;&#x60;. | [optional] 
**QueueBindingName** | Pointer to **string** | The name of a queue in the Message VPN. | [optional] 
**RestDeliveryPointName** | Pointer to **string** | The name of the REST Delivery Point. | [optional] 

## Methods

### NewMsgVpnRestDeliveryPointQueueBinding

`func NewMsgVpnRestDeliveryPointQueueBinding() *MsgVpnRestDeliveryPointQueueBinding`

NewMsgVpnRestDeliveryPointQueueBinding instantiates a new MsgVpnRestDeliveryPointQueueBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnRestDeliveryPointQueueBindingWithDefaults

`func NewMsgVpnRestDeliveryPointQueueBindingWithDefaults() *MsgVpnRestDeliveryPointQueueBinding`

NewMsgVpnRestDeliveryPointQueueBindingWithDefaults instantiates a new MsgVpnRestDeliveryPointQueueBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayReplaceTargetAuthorityEnabled

`func (o *MsgVpnRestDeliveryPointQueueBinding) GetGatewayReplaceTargetAuthorityEnabled() bool`

GetGatewayReplaceTargetAuthorityEnabled returns the GatewayReplaceTargetAuthorityEnabled field if non-nil, zero value otherwise.

### GetGatewayReplaceTargetAuthorityEnabledOk

`func (o *MsgVpnRestDeliveryPointQueueBinding) GetGatewayReplaceTargetAuthorityEnabledOk() (*bool, bool)`

GetGatewayReplaceTargetAuthorityEnabledOk returns a tuple with the GatewayReplaceTargetAuthorityEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayReplaceTargetAuthorityEnabled

`func (o *MsgVpnRestDeliveryPointQueueBinding) SetGatewayReplaceTargetAuthorityEnabled(v bool)`

SetGatewayReplaceTargetAuthorityEnabled sets GatewayReplaceTargetAuthorityEnabled field to given value.

### HasGatewayReplaceTargetAuthorityEnabled

`func (o *MsgVpnRestDeliveryPointQueueBinding) HasGatewayReplaceTargetAuthorityEnabled() bool`

HasGatewayReplaceTargetAuthorityEnabled returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnRestDeliveryPointQueueBinding) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnRestDeliveryPointQueueBinding) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnRestDeliveryPointQueueBinding) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnRestDeliveryPointQueueBinding) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetPostRequestTarget

`func (o *MsgVpnRestDeliveryPointQueueBinding) GetPostRequestTarget() string`

GetPostRequestTarget returns the PostRequestTarget field if non-nil, zero value otherwise.

### GetPostRequestTargetOk

`func (o *MsgVpnRestDeliveryPointQueueBinding) GetPostRequestTargetOk() (*string, bool)`

GetPostRequestTargetOk returns a tuple with the PostRequestTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostRequestTarget

`func (o *MsgVpnRestDeliveryPointQueueBinding) SetPostRequestTarget(v string)`

SetPostRequestTarget sets PostRequestTarget field to given value.

### HasPostRequestTarget

`func (o *MsgVpnRestDeliveryPointQueueBinding) HasPostRequestTarget() bool`

HasPostRequestTarget returns a boolean if a field has been set.

### GetQueueBindingName

`func (o *MsgVpnRestDeliveryPointQueueBinding) GetQueueBindingName() string`

GetQueueBindingName returns the QueueBindingName field if non-nil, zero value otherwise.

### GetQueueBindingNameOk

`func (o *MsgVpnRestDeliveryPointQueueBinding) GetQueueBindingNameOk() (*string, bool)`

GetQueueBindingNameOk returns a tuple with the QueueBindingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueBindingName

`func (o *MsgVpnRestDeliveryPointQueueBinding) SetQueueBindingName(v string)`

SetQueueBindingName sets QueueBindingName field to given value.

### HasQueueBindingName

`func (o *MsgVpnRestDeliveryPointQueueBinding) HasQueueBindingName() bool`

HasQueueBindingName returns a boolean if a field has been set.

### GetRestDeliveryPointName

`func (o *MsgVpnRestDeliveryPointQueueBinding) GetRestDeliveryPointName() string`

GetRestDeliveryPointName returns the RestDeliveryPointName field if non-nil, zero value otherwise.

### GetRestDeliveryPointNameOk

`func (o *MsgVpnRestDeliveryPointQueueBinding) GetRestDeliveryPointNameOk() (*string, bool)`

GetRestDeliveryPointNameOk returns a tuple with the RestDeliveryPointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestDeliveryPointName

`func (o *MsgVpnRestDeliveryPointQueueBinding) SetRestDeliveryPointName(v string)`

SetRestDeliveryPointName sets RestDeliveryPointName field to given value.

### HasRestDeliveryPointName

`func (o *MsgVpnRestDeliveryPointQueueBinding) HasRestDeliveryPointName() bool`

HasRestDeliveryPointName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


