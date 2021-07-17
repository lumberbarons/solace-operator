# MsgVpnReplayLogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to [**MsgVpnReplayLogCollections**](MsgVpnReplayLogCollections.md) |  | [optional] 
**Data** | Pointer to [**MsgVpnReplayLog**](MsgVpnReplayLog.md) |  | [optional] 
**Links** | Pointer to [**MsgVpnReplayLogLinks**](MsgVpnReplayLogLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnReplayLogResponse

`func NewMsgVpnReplayLogResponse(meta SempMeta, ) *MsgVpnReplayLogResponse`

NewMsgVpnReplayLogResponse instantiates a new MsgVpnReplayLogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnReplayLogResponseWithDefaults

`func NewMsgVpnReplayLogResponseWithDefaults() *MsgVpnReplayLogResponse`

NewMsgVpnReplayLogResponseWithDefaults instantiates a new MsgVpnReplayLogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnReplayLogResponse) GetCollections() MsgVpnReplayLogCollections`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnReplayLogResponse) GetCollectionsOk() (*MsgVpnReplayLogCollections, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnReplayLogResponse) SetCollections(v MsgVpnReplayLogCollections)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnReplayLogResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnReplayLogResponse) GetData() MsgVpnReplayLog`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnReplayLogResponse) GetDataOk() (*MsgVpnReplayLog, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnReplayLogResponse) SetData(v MsgVpnReplayLog)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnReplayLogResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnReplayLogResponse) GetLinks() MsgVpnReplayLogLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnReplayLogResponse) GetLinksOk() (*MsgVpnReplayLogLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnReplayLogResponse) SetLinks(v MsgVpnReplayLogLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnReplayLogResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnReplayLogResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnReplayLogResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnReplayLogResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


