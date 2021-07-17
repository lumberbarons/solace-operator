# MsgVpnClientResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to [**MsgVpnClientCollections**](MsgVpnClientCollections.md) |  | [optional] 
**Data** | Pointer to [**MsgVpnClient**](MsgVpnClient.md) |  | [optional] 
**Links** | Pointer to [**MsgVpnClientLinks**](MsgVpnClientLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnClientResponse

`func NewMsgVpnClientResponse(meta SempMeta, ) *MsgVpnClientResponse`

NewMsgVpnClientResponse instantiates a new MsgVpnClientResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnClientResponseWithDefaults

`func NewMsgVpnClientResponseWithDefaults() *MsgVpnClientResponse`

NewMsgVpnClientResponseWithDefaults instantiates a new MsgVpnClientResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnClientResponse) GetCollections() MsgVpnClientCollections`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnClientResponse) GetCollectionsOk() (*MsgVpnClientCollections, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnClientResponse) SetCollections(v MsgVpnClientCollections)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnClientResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnClientResponse) GetData() MsgVpnClient`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnClientResponse) GetDataOk() (*MsgVpnClient, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnClientResponse) SetData(v MsgVpnClient)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnClientResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnClientResponse) GetLinks() MsgVpnClientLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnClientResponse) GetLinksOk() (*MsgVpnClientLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnClientResponse) SetLinks(v MsgVpnClientLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnClientResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnClientResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnClientResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnClientResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


