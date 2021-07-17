# MsgVpnClientTransactedSessionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**[]MsgVpnClientTransactedSession**](MsgVpnClientTransactedSession.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnClientTransactedSessionLinks**](MsgVpnClientTransactedSessionLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnClientTransactedSessionsResponse

`func NewMsgVpnClientTransactedSessionsResponse(meta SempMeta, ) *MsgVpnClientTransactedSessionsResponse`

NewMsgVpnClientTransactedSessionsResponse instantiates a new MsgVpnClientTransactedSessionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnClientTransactedSessionsResponseWithDefaults

`func NewMsgVpnClientTransactedSessionsResponseWithDefaults() *MsgVpnClientTransactedSessionsResponse`

NewMsgVpnClientTransactedSessionsResponseWithDefaults instantiates a new MsgVpnClientTransactedSessionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnClientTransactedSessionsResponse) GetCollections() []map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnClientTransactedSessionsResponse) GetCollectionsOk() (*[]map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnClientTransactedSessionsResponse) SetCollections(v []map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnClientTransactedSessionsResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnClientTransactedSessionsResponse) GetData() []MsgVpnClientTransactedSession`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnClientTransactedSessionsResponse) GetDataOk() (*[]MsgVpnClientTransactedSession, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnClientTransactedSessionsResponse) SetData(v []MsgVpnClientTransactedSession)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnClientTransactedSessionsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnClientTransactedSessionsResponse) GetLinks() []MsgVpnClientTransactedSessionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnClientTransactedSessionsResponse) GetLinksOk() (*[]MsgVpnClientTransactedSessionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnClientTransactedSessionsResponse) SetLinks(v []MsgVpnClientTransactedSessionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnClientTransactedSessionsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnClientTransactedSessionsResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnClientTransactedSessionsResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnClientTransactedSessionsResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


