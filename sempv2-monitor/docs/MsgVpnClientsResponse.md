# MsgVpnClientsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to [**[]MsgVpnClientCollections**](MsgVpnClientCollections.md) |  | [optional] 
**Data** | Pointer to [**[]MsgVpnClient**](MsgVpnClient.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnClientLinks**](MsgVpnClientLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnClientsResponse

`func NewMsgVpnClientsResponse(meta SempMeta, ) *MsgVpnClientsResponse`

NewMsgVpnClientsResponse instantiates a new MsgVpnClientsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnClientsResponseWithDefaults

`func NewMsgVpnClientsResponseWithDefaults() *MsgVpnClientsResponse`

NewMsgVpnClientsResponseWithDefaults instantiates a new MsgVpnClientsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnClientsResponse) GetCollections() []MsgVpnClientCollections`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnClientsResponse) GetCollectionsOk() (*[]MsgVpnClientCollections, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnClientsResponse) SetCollections(v []MsgVpnClientCollections)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnClientsResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnClientsResponse) GetData() []MsgVpnClient`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnClientsResponse) GetDataOk() (*[]MsgVpnClient, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnClientsResponse) SetData(v []MsgVpnClient)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnClientsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnClientsResponse) GetLinks() []MsgVpnClientLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnClientsResponse) GetLinksOk() (*[]MsgVpnClientLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnClientsResponse) SetLinks(v []MsgVpnClientLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnClientsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnClientsResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnClientsResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnClientsResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


