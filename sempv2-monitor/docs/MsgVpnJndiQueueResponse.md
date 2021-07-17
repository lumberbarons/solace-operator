# MsgVpnJndiQueueResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**MsgVpnJndiQueue**](MsgVpnJndiQueue.md) |  | [optional] 
**Links** | Pointer to [**MsgVpnJndiQueueLinks**](MsgVpnJndiQueueLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnJndiQueueResponse

`func NewMsgVpnJndiQueueResponse(meta SempMeta, ) *MsgVpnJndiQueueResponse`

NewMsgVpnJndiQueueResponse instantiates a new MsgVpnJndiQueueResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnJndiQueueResponseWithDefaults

`func NewMsgVpnJndiQueueResponseWithDefaults() *MsgVpnJndiQueueResponse`

NewMsgVpnJndiQueueResponseWithDefaults instantiates a new MsgVpnJndiQueueResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnJndiQueueResponse) GetCollections() map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnJndiQueueResponse) GetCollectionsOk() (*map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnJndiQueueResponse) SetCollections(v map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnJndiQueueResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnJndiQueueResponse) GetData() MsgVpnJndiQueue`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnJndiQueueResponse) GetDataOk() (*MsgVpnJndiQueue, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnJndiQueueResponse) SetData(v MsgVpnJndiQueue)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnJndiQueueResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnJndiQueueResponse) GetLinks() MsgVpnJndiQueueLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnJndiQueueResponse) GetLinksOk() (*MsgVpnJndiQueueLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnJndiQueueResponse) SetLinks(v MsgVpnJndiQueueLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnJndiQueueResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnJndiQueueResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnJndiQueueResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnJndiQueueResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


