# MsgVpnReplayLogMsgsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**[]MsgVpnReplayLogMsg**](MsgVpnReplayLogMsg.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnReplayLogMsgLinks**](MsgVpnReplayLogMsgLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnReplayLogMsgsResponse

`func NewMsgVpnReplayLogMsgsResponse(meta SempMeta, ) *MsgVpnReplayLogMsgsResponse`

NewMsgVpnReplayLogMsgsResponse instantiates a new MsgVpnReplayLogMsgsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnReplayLogMsgsResponseWithDefaults

`func NewMsgVpnReplayLogMsgsResponseWithDefaults() *MsgVpnReplayLogMsgsResponse`

NewMsgVpnReplayLogMsgsResponseWithDefaults instantiates a new MsgVpnReplayLogMsgsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnReplayLogMsgsResponse) GetCollections() []map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnReplayLogMsgsResponse) GetCollectionsOk() (*[]map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnReplayLogMsgsResponse) SetCollections(v []map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnReplayLogMsgsResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnReplayLogMsgsResponse) GetData() []MsgVpnReplayLogMsg`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnReplayLogMsgsResponse) GetDataOk() (*[]MsgVpnReplayLogMsg, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnReplayLogMsgsResponse) SetData(v []MsgVpnReplayLogMsg)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnReplayLogMsgsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnReplayLogMsgsResponse) GetLinks() []MsgVpnReplayLogMsgLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnReplayLogMsgsResponse) GetLinksOk() (*[]MsgVpnReplayLogMsgLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnReplayLogMsgsResponse) SetLinks(v []MsgVpnReplayLogMsgLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnReplayLogMsgsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnReplayLogMsgsResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnReplayLogMsgsResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnReplayLogMsgsResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


