# MsgVpnClientTransactedSessionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**MsgVpnClientTransactedSession**](MsgVpnClientTransactedSession.md) |  | [optional] 
**Links** | Pointer to [**MsgVpnClientTransactedSessionLinks**](MsgVpnClientTransactedSessionLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnClientTransactedSessionResponse

`func NewMsgVpnClientTransactedSessionResponse(meta SempMeta, ) *MsgVpnClientTransactedSessionResponse`

NewMsgVpnClientTransactedSessionResponse instantiates a new MsgVpnClientTransactedSessionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnClientTransactedSessionResponseWithDefaults

`func NewMsgVpnClientTransactedSessionResponseWithDefaults() *MsgVpnClientTransactedSessionResponse`

NewMsgVpnClientTransactedSessionResponseWithDefaults instantiates a new MsgVpnClientTransactedSessionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnClientTransactedSessionResponse) GetCollections() map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnClientTransactedSessionResponse) GetCollectionsOk() (*map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnClientTransactedSessionResponse) SetCollections(v map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnClientTransactedSessionResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnClientTransactedSessionResponse) GetData() MsgVpnClientTransactedSession`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnClientTransactedSessionResponse) GetDataOk() (*MsgVpnClientTransactedSession, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnClientTransactedSessionResponse) SetData(v MsgVpnClientTransactedSession)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnClientTransactedSessionResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnClientTransactedSessionResponse) GetLinks() MsgVpnClientTransactedSessionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnClientTransactedSessionResponse) GetLinksOk() (*MsgVpnClientTransactedSessionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnClientTransactedSessionResponse) SetLinks(v MsgVpnClientTransactedSessionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnClientTransactedSessionResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnClientTransactedSessionResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnClientTransactedSessionResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnClientTransactedSessionResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


