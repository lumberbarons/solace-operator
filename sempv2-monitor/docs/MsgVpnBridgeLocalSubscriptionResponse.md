# MsgVpnBridgeLocalSubscriptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**MsgVpnBridgeLocalSubscription**](MsgVpnBridgeLocalSubscription.md) |  | [optional] 
**Links** | Pointer to [**MsgVpnBridgeLocalSubscriptionLinks**](MsgVpnBridgeLocalSubscriptionLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnBridgeLocalSubscriptionResponse

`func NewMsgVpnBridgeLocalSubscriptionResponse(meta SempMeta, ) *MsgVpnBridgeLocalSubscriptionResponse`

NewMsgVpnBridgeLocalSubscriptionResponse instantiates a new MsgVpnBridgeLocalSubscriptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnBridgeLocalSubscriptionResponseWithDefaults

`func NewMsgVpnBridgeLocalSubscriptionResponseWithDefaults() *MsgVpnBridgeLocalSubscriptionResponse`

NewMsgVpnBridgeLocalSubscriptionResponseWithDefaults instantiates a new MsgVpnBridgeLocalSubscriptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnBridgeLocalSubscriptionResponse) GetCollections() map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnBridgeLocalSubscriptionResponse) GetCollectionsOk() (*map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnBridgeLocalSubscriptionResponse) SetCollections(v map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnBridgeLocalSubscriptionResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnBridgeLocalSubscriptionResponse) GetData() MsgVpnBridgeLocalSubscription`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnBridgeLocalSubscriptionResponse) GetDataOk() (*MsgVpnBridgeLocalSubscription, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnBridgeLocalSubscriptionResponse) SetData(v MsgVpnBridgeLocalSubscription)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnBridgeLocalSubscriptionResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnBridgeLocalSubscriptionResponse) GetLinks() MsgVpnBridgeLocalSubscriptionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnBridgeLocalSubscriptionResponse) GetLinksOk() (*MsgVpnBridgeLocalSubscriptionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnBridgeLocalSubscriptionResponse) SetLinks(v MsgVpnBridgeLocalSubscriptionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnBridgeLocalSubscriptionResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnBridgeLocalSubscriptionResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnBridgeLocalSubscriptionResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnBridgeLocalSubscriptionResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


