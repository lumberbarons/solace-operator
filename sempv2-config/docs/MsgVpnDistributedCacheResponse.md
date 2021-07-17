# MsgVpnDistributedCacheResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**MsgVpnDistributedCache**](MsgVpnDistributedCache.md) |  | [optional] 
**Links** | Pointer to [**MsgVpnDistributedCacheLinks**](MsgVpnDistributedCacheLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnDistributedCacheResponse

`func NewMsgVpnDistributedCacheResponse(meta SempMeta, ) *MsgVpnDistributedCacheResponse`

NewMsgVpnDistributedCacheResponse instantiates a new MsgVpnDistributedCacheResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnDistributedCacheResponseWithDefaults

`func NewMsgVpnDistributedCacheResponseWithDefaults() *MsgVpnDistributedCacheResponse`

NewMsgVpnDistributedCacheResponseWithDefaults instantiates a new MsgVpnDistributedCacheResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MsgVpnDistributedCacheResponse) GetData() MsgVpnDistributedCache`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnDistributedCacheResponse) GetDataOk() (*MsgVpnDistributedCache, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnDistributedCacheResponse) SetData(v MsgVpnDistributedCache)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnDistributedCacheResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnDistributedCacheResponse) GetLinks() MsgVpnDistributedCacheLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnDistributedCacheResponse) GetLinksOk() (*MsgVpnDistributedCacheLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnDistributedCacheResponse) SetLinks(v MsgVpnDistributedCacheLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnDistributedCacheResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnDistributedCacheResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnDistributedCacheResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnDistributedCacheResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


