# MsgVpnReplayLogMsgResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**MsgVpnReplayLogMsg**](MsgVpnReplayLogMsg.md) |  | [optional] 
**Links** | Pointer to [**MsgVpnReplayLogMsgLinks**](MsgVpnReplayLogMsgLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnReplayLogMsgResponse

`func NewMsgVpnReplayLogMsgResponse(meta SempMeta, ) *MsgVpnReplayLogMsgResponse`

NewMsgVpnReplayLogMsgResponse instantiates a new MsgVpnReplayLogMsgResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnReplayLogMsgResponseWithDefaults

`func NewMsgVpnReplayLogMsgResponseWithDefaults() *MsgVpnReplayLogMsgResponse`

NewMsgVpnReplayLogMsgResponseWithDefaults instantiates a new MsgVpnReplayLogMsgResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnReplayLogMsgResponse) GetCollections() map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnReplayLogMsgResponse) GetCollectionsOk() (*map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnReplayLogMsgResponse) SetCollections(v map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnReplayLogMsgResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnReplayLogMsgResponse) GetData() MsgVpnReplayLogMsg`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnReplayLogMsgResponse) GetDataOk() (*MsgVpnReplayLogMsg, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnReplayLogMsgResponse) SetData(v MsgVpnReplayLogMsg)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnReplayLogMsgResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnReplayLogMsgResponse) GetLinks() MsgVpnReplayLogMsgLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnReplayLogMsgResponse) GetLinksOk() (*MsgVpnReplayLogMsgLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnReplayLogMsgResponse) SetLinks(v MsgVpnReplayLogMsgLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnReplayLogMsgResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnReplayLogMsgResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnReplayLogMsgResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnReplayLogMsgResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


