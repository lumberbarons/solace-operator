# MsgVpnClientRxFlowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**MsgVpnClientRxFlow**](MsgVpnClientRxFlow.md) |  | [optional] 
**Links** | Pointer to [**MsgVpnClientRxFlowLinks**](MsgVpnClientRxFlowLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnClientRxFlowResponse

`func NewMsgVpnClientRxFlowResponse(meta SempMeta, ) *MsgVpnClientRxFlowResponse`

NewMsgVpnClientRxFlowResponse instantiates a new MsgVpnClientRxFlowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnClientRxFlowResponseWithDefaults

`func NewMsgVpnClientRxFlowResponseWithDefaults() *MsgVpnClientRxFlowResponse`

NewMsgVpnClientRxFlowResponseWithDefaults instantiates a new MsgVpnClientRxFlowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnClientRxFlowResponse) GetCollections() map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnClientRxFlowResponse) GetCollectionsOk() (*map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnClientRxFlowResponse) SetCollections(v map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnClientRxFlowResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnClientRxFlowResponse) GetData() MsgVpnClientRxFlow`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnClientRxFlowResponse) GetDataOk() (*MsgVpnClientRxFlow, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnClientRxFlowResponse) SetData(v MsgVpnClientRxFlow)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnClientRxFlowResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnClientRxFlowResponse) GetLinks() MsgVpnClientRxFlowLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnClientRxFlowResponse) GetLinksOk() (*MsgVpnClientRxFlowLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnClientRxFlowResponse) SetLinks(v MsgVpnClientRxFlowLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnClientRxFlowResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnClientRxFlowResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnClientRxFlowResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnClientRxFlowResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


