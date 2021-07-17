# MsgVpnDistributedCacheClusterInstancesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to [**[]MsgVpnDistributedCacheClusterInstanceCollections**](MsgVpnDistributedCacheClusterInstanceCollections.md) |  | [optional] 
**Data** | Pointer to [**[]MsgVpnDistributedCacheClusterInstance**](MsgVpnDistributedCacheClusterInstance.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnDistributedCacheClusterInstanceLinks**](MsgVpnDistributedCacheClusterInstanceLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnDistributedCacheClusterInstancesResponse

`func NewMsgVpnDistributedCacheClusterInstancesResponse(meta SempMeta, ) *MsgVpnDistributedCacheClusterInstancesResponse`

NewMsgVpnDistributedCacheClusterInstancesResponse instantiates a new MsgVpnDistributedCacheClusterInstancesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnDistributedCacheClusterInstancesResponseWithDefaults

`func NewMsgVpnDistributedCacheClusterInstancesResponseWithDefaults() *MsgVpnDistributedCacheClusterInstancesResponse`

NewMsgVpnDistributedCacheClusterInstancesResponseWithDefaults instantiates a new MsgVpnDistributedCacheClusterInstancesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnDistributedCacheClusterInstancesResponse) GetCollections() []MsgVpnDistributedCacheClusterInstanceCollections`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnDistributedCacheClusterInstancesResponse) GetCollectionsOk() (*[]MsgVpnDistributedCacheClusterInstanceCollections, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnDistributedCacheClusterInstancesResponse) SetCollections(v []MsgVpnDistributedCacheClusterInstanceCollections)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnDistributedCacheClusterInstancesResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnDistributedCacheClusterInstancesResponse) GetData() []MsgVpnDistributedCacheClusterInstance`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnDistributedCacheClusterInstancesResponse) GetDataOk() (*[]MsgVpnDistributedCacheClusterInstance, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnDistributedCacheClusterInstancesResponse) SetData(v []MsgVpnDistributedCacheClusterInstance)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnDistributedCacheClusterInstancesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnDistributedCacheClusterInstancesResponse) GetLinks() []MsgVpnDistributedCacheClusterInstanceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnDistributedCacheClusterInstancesResponse) GetLinksOk() (*[]MsgVpnDistributedCacheClusterInstanceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnDistributedCacheClusterInstancesResponse) SetLinks(v []MsgVpnDistributedCacheClusterInstanceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnDistributedCacheClusterInstancesResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnDistributedCacheClusterInstancesResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnDistributedCacheClusterInstancesResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnDistributedCacheClusterInstancesResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


