# MsgVpnReplayLogsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to [**[]MsgVpnReplayLogCollections**](MsgVpnReplayLogCollections.md) |  | [optional] 
**Data** | Pointer to [**[]MsgVpnReplayLog**](MsgVpnReplayLog.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnReplayLogLinks**](MsgVpnReplayLogLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnReplayLogsResponse

`func NewMsgVpnReplayLogsResponse(meta SempMeta, ) *MsgVpnReplayLogsResponse`

NewMsgVpnReplayLogsResponse instantiates a new MsgVpnReplayLogsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnReplayLogsResponseWithDefaults

`func NewMsgVpnReplayLogsResponseWithDefaults() *MsgVpnReplayLogsResponse`

NewMsgVpnReplayLogsResponseWithDefaults instantiates a new MsgVpnReplayLogsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnReplayLogsResponse) GetCollections() []MsgVpnReplayLogCollections`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnReplayLogsResponse) GetCollectionsOk() (*[]MsgVpnReplayLogCollections, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnReplayLogsResponse) SetCollections(v []MsgVpnReplayLogCollections)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnReplayLogsResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnReplayLogsResponse) GetData() []MsgVpnReplayLog`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnReplayLogsResponse) GetDataOk() (*[]MsgVpnReplayLog, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnReplayLogsResponse) SetData(v []MsgVpnReplayLog)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnReplayLogsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnReplayLogsResponse) GetLinks() []MsgVpnReplayLogLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnReplayLogsResponse) GetLinksOk() (*[]MsgVpnReplayLogLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnReplayLogsResponse) SetLinks(v []MsgVpnReplayLogLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnReplayLogsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnReplayLogsResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnReplayLogsResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnReplayLogsResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


