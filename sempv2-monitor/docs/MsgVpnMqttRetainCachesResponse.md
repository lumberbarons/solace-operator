# MsgVpnMqttRetainCachesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**[]MsgVpnMqttRetainCache**](MsgVpnMqttRetainCache.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnMqttRetainCacheLinks**](MsgVpnMqttRetainCacheLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnMqttRetainCachesResponse

`func NewMsgVpnMqttRetainCachesResponse(meta SempMeta, ) *MsgVpnMqttRetainCachesResponse`

NewMsgVpnMqttRetainCachesResponse instantiates a new MsgVpnMqttRetainCachesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnMqttRetainCachesResponseWithDefaults

`func NewMsgVpnMqttRetainCachesResponseWithDefaults() *MsgVpnMqttRetainCachesResponse`

NewMsgVpnMqttRetainCachesResponseWithDefaults instantiates a new MsgVpnMqttRetainCachesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnMqttRetainCachesResponse) GetCollections() []map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnMqttRetainCachesResponse) GetCollectionsOk() (*[]map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnMqttRetainCachesResponse) SetCollections(v []map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnMqttRetainCachesResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnMqttRetainCachesResponse) GetData() []MsgVpnMqttRetainCache`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnMqttRetainCachesResponse) GetDataOk() (*[]MsgVpnMqttRetainCache, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnMqttRetainCachesResponse) SetData(v []MsgVpnMqttRetainCache)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnMqttRetainCachesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnMqttRetainCachesResponse) GetLinks() []MsgVpnMqttRetainCacheLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnMqttRetainCachesResponse) GetLinksOk() (*[]MsgVpnMqttRetainCacheLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnMqttRetainCachesResponse) SetLinks(v []MsgVpnMqttRetainCacheLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnMqttRetainCachesResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnMqttRetainCachesResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnMqttRetainCachesResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnMqttRetainCachesResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


