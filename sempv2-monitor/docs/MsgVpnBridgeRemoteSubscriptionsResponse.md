# MsgVpnBridgeRemoteSubscriptionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**[]MsgVpnBridgeRemoteSubscription**](MsgVpnBridgeRemoteSubscription.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnBridgeRemoteSubscriptionLinks**](MsgVpnBridgeRemoteSubscriptionLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnBridgeRemoteSubscriptionsResponse

`func NewMsgVpnBridgeRemoteSubscriptionsResponse(meta SempMeta, ) *MsgVpnBridgeRemoteSubscriptionsResponse`

NewMsgVpnBridgeRemoteSubscriptionsResponse instantiates a new MsgVpnBridgeRemoteSubscriptionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnBridgeRemoteSubscriptionsResponseWithDefaults

`func NewMsgVpnBridgeRemoteSubscriptionsResponseWithDefaults() *MsgVpnBridgeRemoteSubscriptionsResponse`

NewMsgVpnBridgeRemoteSubscriptionsResponseWithDefaults instantiates a new MsgVpnBridgeRemoteSubscriptionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnBridgeRemoteSubscriptionsResponse) GetCollections() []map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnBridgeRemoteSubscriptionsResponse) GetCollectionsOk() (*[]map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnBridgeRemoteSubscriptionsResponse) SetCollections(v []map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnBridgeRemoteSubscriptionsResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnBridgeRemoteSubscriptionsResponse) GetData() []MsgVpnBridgeRemoteSubscription`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnBridgeRemoteSubscriptionsResponse) GetDataOk() (*[]MsgVpnBridgeRemoteSubscription, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnBridgeRemoteSubscriptionsResponse) SetData(v []MsgVpnBridgeRemoteSubscription)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnBridgeRemoteSubscriptionsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnBridgeRemoteSubscriptionsResponse) GetLinks() []MsgVpnBridgeRemoteSubscriptionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnBridgeRemoteSubscriptionsResponse) GetLinksOk() (*[]MsgVpnBridgeRemoteSubscriptionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnBridgeRemoteSubscriptionsResponse) SetLinks(v []MsgVpnBridgeRemoteSubscriptionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnBridgeRemoteSubscriptionsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnBridgeRemoteSubscriptionsResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnBridgeRemoteSubscriptionsResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnBridgeRemoteSubscriptionsResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


