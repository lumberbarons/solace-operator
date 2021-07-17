# MsgVpnClientConnectionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**[]MsgVpnClientConnection**](MsgVpnClientConnection.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnClientConnectionLinks**](MsgVpnClientConnectionLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnClientConnectionsResponse

`func NewMsgVpnClientConnectionsResponse(meta SempMeta, ) *MsgVpnClientConnectionsResponse`

NewMsgVpnClientConnectionsResponse instantiates a new MsgVpnClientConnectionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnClientConnectionsResponseWithDefaults

`func NewMsgVpnClientConnectionsResponseWithDefaults() *MsgVpnClientConnectionsResponse`

NewMsgVpnClientConnectionsResponseWithDefaults instantiates a new MsgVpnClientConnectionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnClientConnectionsResponse) GetCollections() []map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnClientConnectionsResponse) GetCollectionsOk() (*[]map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnClientConnectionsResponse) SetCollections(v []map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnClientConnectionsResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnClientConnectionsResponse) GetData() []MsgVpnClientConnection`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnClientConnectionsResponse) GetDataOk() (*[]MsgVpnClientConnection, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnClientConnectionsResponse) SetData(v []MsgVpnClientConnection)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnClientConnectionsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnClientConnectionsResponse) GetLinks() []MsgVpnClientConnectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnClientConnectionsResponse) GetLinksOk() (*[]MsgVpnClientConnectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnClientConnectionsResponse) SetLinks(v []MsgVpnClientConnectionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnClientConnectionsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnClientConnectionsResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnClientConnectionsResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnClientConnectionsResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


