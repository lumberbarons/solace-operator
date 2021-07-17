# MsgVpnBridgesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to [**[]MsgVpnBridgeCollections**](MsgVpnBridgeCollections.md) |  | [optional] 
**Data** | Pointer to [**[]MsgVpnBridge**](MsgVpnBridge.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnBridgeLinks**](MsgVpnBridgeLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnBridgesResponse

`func NewMsgVpnBridgesResponse(meta SempMeta, ) *MsgVpnBridgesResponse`

NewMsgVpnBridgesResponse instantiates a new MsgVpnBridgesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnBridgesResponseWithDefaults

`func NewMsgVpnBridgesResponseWithDefaults() *MsgVpnBridgesResponse`

NewMsgVpnBridgesResponseWithDefaults instantiates a new MsgVpnBridgesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnBridgesResponse) GetCollections() []MsgVpnBridgeCollections`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnBridgesResponse) GetCollectionsOk() (*[]MsgVpnBridgeCollections, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnBridgesResponse) SetCollections(v []MsgVpnBridgeCollections)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnBridgesResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnBridgesResponse) GetData() []MsgVpnBridge`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnBridgesResponse) GetDataOk() (*[]MsgVpnBridge, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnBridgesResponse) SetData(v []MsgVpnBridge)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnBridgesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnBridgesResponse) GetLinks() []MsgVpnBridgeLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnBridgesResponse) GetLinksOk() (*[]MsgVpnBridgeLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnBridgesResponse) SetLinks(v []MsgVpnBridgeLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnBridgesResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnBridgesResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnBridgesResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnBridgesResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


