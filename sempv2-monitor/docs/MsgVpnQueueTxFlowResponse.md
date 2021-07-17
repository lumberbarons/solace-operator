# MsgVpnQueueTxFlowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**MsgVpnQueueTxFlow**](MsgVpnQueueTxFlow.md) |  | [optional] 
**Links** | Pointer to [**MsgVpnQueueTxFlowLinks**](MsgVpnQueueTxFlowLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnQueueTxFlowResponse

`func NewMsgVpnQueueTxFlowResponse(meta SempMeta, ) *MsgVpnQueueTxFlowResponse`

NewMsgVpnQueueTxFlowResponse instantiates a new MsgVpnQueueTxFlowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnQueueTxFlowResponseWithDefaults

`func NewMsgVpnQueueTxFlowResponseWithDefaults() *MsgVpnQueueTxFlowResponse`

NewMsgVpnQueueTxFlowResponseWithDefaults instantiates a new MsgVpnQueueTxFlowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnQueueTxFlowResponse) GetCollections() map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnQueueTxFlowResponse) GetCollectionsOk() (*map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnQueueTxFlowResponse) SetCollections(v map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnQueueTxFlowResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnQueueTxFlowResponse) GetData() MsgVpnQueueTxFlow`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnQueueTxFlowResponse) GetDataOk() (*MsgVpnQueueTxFlow, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnQueueTxFlowResponse) SetData(v MsgVpnQueueTxFlow)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnQueueTxFlowResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnQueueTxFlowResponse) GetLinks() MsgVpnQueueTxFlowLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnQueueTxFlowResponse) GetLinksOk() (*MsgVpnQueueTxFlowLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnQueueTxFlowResponse) SetLinks(v MsgVpnQueueTxFlowLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnQueueTxFlowResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnQueueTxFlowResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnQueueTxFlowResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnQueueTxFlowResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


