# MsgVpnQueueMsgsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**[]MsgVpnQueueMsg**](MsgVpnQueueMsg.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnQueueMsgLinks**](MsgVpnQueueMsgLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnQueueMsgsResponse

`func NewMsgVpnQueueMsgsResponse(meta SempMeta, ) *MsgVpnQueueMsgsResponse`

NewMsgVpnQueueMsgsResponse instantiates a new MsgVpnQueueMsgsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnQueueMsgsResponseWithDefaults

`func NewMsgVpnQueueMsgsResponseWithDefaults() *MsgVpnQueueMsgsResponse`

NewMsgVpnQueueMsgsResponseWithDefaults instantiates a new MsgVpnQueueMsgsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnQueueMsgsResponse) GetCollections() []map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnQueueMsgsResponse) GetCollectionsOk() (*[]map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnQueueMsgsResponse) SetCollections(v []map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnQueueMsgsResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnQueueMsgsResponse) GetData() []MsgVpnQueueMsg`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnQueueMsgsResponse) GetDataOk() (*[]MsgVpnQueueMsg, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnQueueMsgsResponse) SetData(v []MsgVpnQueueMsg)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnQueueMsgsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnQueueMsgsResponse) GetLinks() []MsgVpnQueueMsgLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnQueueMsgsResponse) GetLinksOk() (*[]MsgVpnQueueMsgLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnQueueMsgsResponse) SetLinks(v []MsgVpnQueueMsgLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnQueueMsgsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnQueueMsgsResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnQueueMsgsResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnQueueMsgsResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


