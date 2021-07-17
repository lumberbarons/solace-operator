# MsgVpnQueueResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**MsgVpnQueue**](MsgVpnQueue.md) |  | [optional] 
**Links** | Pointer to [**MsgVpnQueueLinks**](MsgVpnQueueLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnQueueResponse

`func NewMsgVpnQueueResponse(meta SempMeta, ) *MsgVpnQueueResponse`

NewMsgVpnQueueResponse instantiates a new MsgVpnQueueResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnQueueResponseWithDefaults

`func NewMsgVpnQueueResponseWithDefaults() *MsgVpnQueueResponse`

NewMsgVpnQueueResponseWithDefaults instantiates a new MsgVpnQueueResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MsgVpnQueueResponse) GetData() MsgVpnQueue`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnQueueResponse) GetDataOk() (*MsgVpnQueue, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnQueueResponse) SetData(v MsgVpnQueue)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnQueueResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnQueueResponse) GetLinks() MsgVpnQueueLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnQueueResponse) GetLinksOk() (*MsgVpnQueueLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnQueueResponse) SetLinks(v MsgVpnQueueLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnQueueResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnQueueResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnQueueResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnQueueResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


