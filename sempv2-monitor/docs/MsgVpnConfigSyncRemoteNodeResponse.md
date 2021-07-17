# MsgVpnConfigSyncRemoteNodeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**MsgVpnConfigSyncRemoteNode**](MsgVpnConfigSyncRemoteNode.md) |  | [optional] 
**Links** | Pointer to [**MsgVpnConfigSyncRemoteNodeLinks**](MsgVpnConfigSyncRemoteNodeLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnConfigSyncRemoteNodeResponse

`func NewMsgVpnConfigSyncRemoteNodeResponse(meta SempMeta, ) *MsgVpnConfigSyncRemoteNodeResponse`

NewMsgVpnConfigSyncRemoteNodeResponse instantiates a new MsgVpnConfigSyncRemoteNodeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnConfigSyncRemoteNodeResponseWithDefaults

`func NewMsgVpnConfigSyncRemoteNodeResponseWithDefaults() *MsgVpnConfigSyncRemoteNodeResponse`

NewMsgVpnConfigSyncRemoteNodeResponseWithDefaults instantiates a new MsgVpnConfigSyncRemoteNodeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnConfigSyncRemoteNodeResponse) GetCollections() map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnConfigSyncRemoteNodeResponse) GetCollectionsOk() (*map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnConfigSyncRemoteNodeResponse) SetCollections(v map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnConfigSyncRemoteNodeResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnConfigSyncRemoteNodeResponse) GetData() MsgVpnConfigSyncRemoteNode`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnConfigSyncRemoteNodeResponse) GetDataOk() (*MsgVpnConfigSyncRemoteNode, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnConfigSyncRemoteNodeResponse) SetData(v MsgVpnConfigSyncRemoteNode)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnConfigSyncRemoteNodeResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnConfigSyncRemoteNodeResponse) GetLinks() MsgVpnConfigSyncRemoteNodeLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnConfigSyncRemoteNodeResponse) GetLinksOk() (*MsgVpnConfigSyncRemoteNodeLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnConfigSyncRemoteNodeResponse) SetLinks(v MsgVpnConfigSyncRemoteNodeLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnConfigSyncRemoteNodeResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnConfigSyncRemoteNodeResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnConfigSyncRemoteNodeResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnConfigSyncRemoteNodeResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


