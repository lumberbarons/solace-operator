# MsgVpnQueueMsgResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**MsgVpnQueueMsg**](MsgVpnQueueMsg.md) |  | [optional] 
**Links** | Pointer to [**MsgVpnQueueMsgLinks**](MsgVpnQueueMsgLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnQueueMsgResponse

`func NewMsgVpnQueueMsgResponse(meta SempMeta, ) *MsgVpnQueueMsgResponse`

NewMsgVpnQueueMsgResponse instantiates a new MsgVpnQueueMsgResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnQueueMsgResponseWithDefaults

`func NewMsgVpnQueueMsgResponseWithDefaults() *MsgVpnQueueMsgResponse`

NewMsgVpnQueueMsgResponseWithDefaults instantiates a new MsgVpnQueueMsgResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnQueueMsgResponse) GetCollections() map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnQueueMsgResponse) GetCollectionsOk() (*map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnQueueMsgResponse) SetCollections(v map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnQueueMsgResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnQueueMsgResponse) GetData() MsgVpnQueueMsg`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnQueueMsgResponse) GetDataOk() (*MsgVpnQueueMsg, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnQueueMsgResponse) SetData(v MsgVpnQueueMsg)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnQueueMsgResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnQueueMsgResponse) GetLinks() MsgVpnQueueMsgLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnQueueMsgResponse) GetLinksOk() (*MsgVpnQueueMsgLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnQueueMsgResponse) SetLinks(v MsgVpnQueueMsgLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnQueueMsgResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnQueueMsgResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnQueueMsgResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnQueueMsgResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


