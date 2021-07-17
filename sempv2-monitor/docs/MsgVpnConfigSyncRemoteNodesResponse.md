# MsgVpnConfigSyncRemoteNodesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**[]MsgVpnConfigSyncRemoteNode**](MsgVpnConfigSyncRemoteNode.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnConfigSyncRemoteNodeLinks**](MsgVpnConfigSyncRemoteNodeLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnConfigSyncRemoteNodesResponse

`func NewMsgVpnConfigSyncRemoteNodesResponse(meta SempMeta, ) *MsgVpnConfigSyncRemoteNodesResponse`

NewMsgVpnConfigSyncRemoteNodesResponse instantiates a new MsgVpnConfigSyncRemoteNodesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnConfigSyncRemoteNodesResponseWithDefaults

`func NewMsgVpnConfigSyncRemoteNodesResponseWithDefaults() *MsgVpnConfigSyncRemoteNodesResponse`

NewMsgVpnConfigSyncRemoteNodesResponseWithDefaults instantiates a new MsgVpnConfigSyncRemoteNodesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnConfigSyncRemoteNodesResponse) GetCollections() []map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnConfigSyncRemoteNodesResponse) GetCollectionsOk() (*[]map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnConfigSyncRemoteNodesResponse) SetCollections(v []map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnConfigSyncRemoteNodesResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnConfigSyncRemoteNodesResponse) GetData() []MsgVpnConfigSyncRemoteNode`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnConfigSyncRemoteNodesResponse) GetDataOk() (*[]MsgVpnConfigSyncRemoteNode, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnConfigSyncRemoteNodesResponse) SetData(v []MsgVpnConfigSyncRemoteNode)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnConfigSyncRemoteNodesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnConfigSyncRemoteNodesResponse) GetLinks() []MsgVpnConfigSyncRemoteNodeLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnConfigSyncRemoteNodesResponse) GetLinksOk() (*[]MsgVpnConfigSyncRemoteNodeLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnConfigSyncRemoteNodesResponse) SetLinks(v []MsgVpnConfigSyncRemoteNodeLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnConfigSyncRemoteNodesResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnConfigSyncRemoteNodesResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnConfigSyncRemoteNodesResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnConfigSyncRemoteNodesResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


