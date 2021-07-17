# MsgVpnTransactionPublisherMsgResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**MsgVpnTransactionPublisherMsg**](MsgVpnTransactionPublisherMsg.md) |  | [optional] 
**Links** | Pointer to [**MsgVpnTransactionPublisherMsgLinks**](MsgVpnTransactionPublisherMsgLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnTransactionPublisherMsgResponse

`func NewMsgVpnTransactionPublisherMsgResponse(meta SempMeta, ) *MsgVpnTransactionPublisherMsgResponse`

NewMsgVpnTransactionPublisherMsgResponse instantiates a new MsgVpnTransactionPublisherMsgResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnTransactionPublisherMsgResponseWithDefaults

`func NewMsgVpnTransactionPublisherMsgResponseWithDefaults() *MsgVpnTransactionPublisherMsgResponse`

NewMsgVpnTransactionPublisherMsgResponseWithDefaults instantiates a new MsgVpnTransactionPublisherMsgResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnTransactionPublisherMsgResponse) GetCollections() map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnTransactionPublisherMsgResponse) GetCollectionsOk() (*map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnTransactionPublisherMsgResponse) SetCollections(v map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnTransactionPublisherMsgResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnTransactionPublisherMsgResponse) GetData() MsgVpnTransactionPublisherMsg`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnTransactionPublisherMsgResponse) GetDataOk() (*MsgVpnTransactionPublisherMsg, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnTransactionPublisherMsgResponse) SetData(v MsgVpnTransactionPublisherMsg)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnTransactionPublisherMsgResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnTransactionPublisherMsgResponse) GetLinks() MsgVpnTransactionPublisherMsgLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnTransactionPublisherMsgResponse) GetLinksOk() (*MsgVpnTransactionPublisherMsgLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnTransactionPublisherMsgResponse) SetLinks(v MsgVpnTransactionPublisherMsgLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnTransactionPublisherMsgResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnTransactionPublisherMsgResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnTransactionPublisherMsgResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnTransactionPublisherMsgResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


