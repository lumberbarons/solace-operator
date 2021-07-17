# MsgVpnClientConnectionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**MsgVpnClientConnection**](MsgVpnClientConnection.md) |  | [optional] 
**Links** | Pointer to [**MsgVpnClientConnectionLinks**](MsgVpnClientConnectionLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnClientConnectionResponse

`func NewMsgVpnClientConnectionResponse(meta SempMeta, ) *MsgVpnClientConnectionResponse`

NewMsgVpnClientConnectionResponse instantiates a new MsgVpnClientConnectionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnClientConnectionResponseWithDefaults

`func NewMsgVpnClientConnectionResponseWithDefaults() *MsgVpnClientConnectionResponse`

NewMsgVpnClientConnectionResponseWithDefaults instantiates a new MsgVpnClientConnectionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnClientConnectionResponse) GetCollections() map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnClientConnectionResponse) GetCollectionsOk() (*map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnClientConnectionResponse) SetCollections(v map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnClientConnectionResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnClientConnectionResponse) GetData() MsgVpnClientConnection`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnClientConnectionResponse) GetDataOk() (*MsgVpnClientConnection, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnClientConnectionResponse) SetData(v MsgVpnClientConnection)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnClientConnectionResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnClientConnectionResponse) GetLinks() MsgVpnClientConnectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnClientConnectionResponse) GetLinksOk() (*MsgVpnClientConnectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnClientConnectionResponse) SetLinks(v MsgVpnClientConnectionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnClientConnectionResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnClientConnectionResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnClientConnectionResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnClientConnectionResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


