# MsgVpnTransactionConsumerMsgResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**MsgVpnTransactionConsumerMsg**](MsgVpnTransactionConsumerMsg.md) |  | [optional] 
**Links** | Pointer to [**MsgVpnTransactionConsumerMsgLinks**](MsgVpnTransactionConsumerMsgLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnTransactionConsumerMsgResponse

`func NewMsgVpnTransactionConsumerMsgResponse(meta SempMeta, ) *MsgVpnTransactionConsumerMsgResponse`

NewMsgVpnTransactionConsumerMsgResponse instantiates a new MsgVpnTransactionConsumerMsgResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnTransactionConsumerMsgResponseWithDefaults

`func NewMsgVpnTransactionConsumerMsgResponseWithDefaults() *MsgVpnTransactionConsumerMsgResponse`

NewMsgVpnTransactionConsumerMsgResponseWithDefaults instantiates a new MsgVpnTransactionConsumerMsgResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnTransactionConsumerMsgResponse) GetCollections() map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnTransactionConsumerMsgResponse) GetCollectionsOk() (*map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnTransactionConsumerMsgResponse) SetCollections(v map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnTransactionConsumerMsgResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnTransactionConsumerMsgResponse) GetData() MsgVpnTransactionConsumerMsg`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnTransactionConsumerMsgResponse) GetDataOk() (*MsgVpnTransactionConsumerMsg, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnTransactionConsumerMsgResponse) SetData(v MsgVpnTransactionConsumerMsg)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnTransactionConsumerMsgResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnTransactionConsumerMsgResponse) GetLinks() MsgVpnTransactionConsumerMsgLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnTransactionConsumerMsgResponse) GetLinksOk() (*MsgVpnTransactionConsumerMsgLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnTransactionConsumerMsgResponse) SetLinks(v MsgVpnTransactionConsumerMsgLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnTransactionConsumerMsgResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnTransactionConsumerMsgResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnTransactionConsumerMsgResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnTransactionConsumerMsgResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


