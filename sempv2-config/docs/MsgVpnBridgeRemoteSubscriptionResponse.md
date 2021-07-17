# MsgVpnBridgeRemoteSubscriptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**MsgVpnBridgeRemoteSubscription**](MsgVpnBridgeRemoteSubscription.md) |  | [optional] 
**Links** | Pointer to [**MsgVpnBridgeRemoteSubscriptionLinks**](MsgVpnBridgeRemoteSubscriptionLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnBridgeRemoteSubscriptionResponse

`func NewMsgVpnBridgeRemoteSubscriptionResponse(meta SempMeta, ) *MsgVpnBridgeRemoteSubscriptionResponse`

NewMsgVpnBridgeRemoteSubscriptionResponse instantiates a new MsgVpnBridgeRemoteSubscriptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnBridgeRemoteSubscriptionResponseWithDefaults

`func NewMsgVpnBridgeRemoteSubscriptionResponseWithDefaults() *MsgVpnBridgeRemoteSubscriptionResponse`

NewMsgVpnBridgeRemoteSubscriptionResponseWithDefaults instantiates a new MsgVpnBridgeRemoteSubscriptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MsgVpnBridgeRemoteSubscriptionResponse) GetData() MsgVpnBridgeRemoteSubscription`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnBridgeRemoteSubscriptionResponse) GetDataOk() (*MsgVpnBridgeRemoteSubscription, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnBridgeRemoteSubscriptionResponse) SetData(v MsgVpnBridgeRemoteSubscription)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnBridgeRemoteSubscriptionResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnBridgeRemoteSubscriptionResponse) GetLinks() MsgVpnBridgeRemoteSubscriptionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnBridgeRemoteSubscriptionResponse) GetLinksOk() (*MsgVpnBridgeRemoteSubscriptionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnBridgeRemoteSubscriptionResponse) SetLinks(v MsgVpnBridgeRemoteSubscriptionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnBridgeRemoteSubscriptionResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnBridgeRemoteSubscriptionResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnBridgeRemoteSubscriptionResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnBridgeRemoteSubscriptionResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


