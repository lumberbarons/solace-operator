# MsgVpnDistributedCachesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]MsgVpnDistributedCache**](MsgVpnDistributedCache.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnDistributedCacheLinks**](MsgVpnDistributedCacheLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnDistributedCachesResponse

`func NewMsgVpnDistributedCachesResponse(meta SempMeta, ) *MsgVpnDistributedCachesResponse`

NewMsgVpnDistributedCachesResponse instantiates a new MsgVpnDistributedCachesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnDistributedCachesResponseWithDefaults

`func NewMsgVpnDistributedCachesResponseWithDefaults() *MsgVpnDistributedCachesResponse`

NewMsgVpnDistributedCachesResponseWithDefaults instantiates a new MsgVpnDistributedCachesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MsgVpnDistributedCachesResponse) GetData() []MsgVpnDistributedCache`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnDistributedCachesResponse) GetDataOk() (*[]MsgVpnDistributedCache, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnDistributedCachesResponse) SetData(v []MsgVpnDistributedCache)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnDistributedCachesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnDistributedCachesResponse) GetLinks() []MsgVpnDistributedCacheLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnDistributedCachesResponse) GetLinksOk() (*[]MsgVpnDistributedCacheLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnDistributedCachesResponse) SetLinks(v []MsgVpnDistributedCacheLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnDistributedCachesResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnDistributedCachesResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnDistributedCachesResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnDistributedCachesResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


