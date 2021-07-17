# MsgVpnDistributedCacheClusterTopicsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]MsgVpnDistributedCacheClusterTopic**](MsgVpnDistributedCacheClusterTopic.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnDistributedCacheClusterTopicLinks**](MsgVpnDistributedCacheClusterTopicLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnDistributedCacheClusterTopicsResponse

`func NewMsgVpnDistributedCacheClusterTopicsResponse(meta SempMeta, ) *MsgVpnDistributedCacheClusterTopicsResponse`

NewMsgVpnDistributedCacheClusterTopicsResponse instantiates a new MsgVpnDistributedCacheClusterTopicsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnDistributedCacheClusterTopicsResponseWithDefaults

`func NewMsgVpnDistributedCacheClusterTopicsResponseWithDefaults() *MsgVpnDistributedCacheClusterTopicsResponse`

NewMsgVpnDistributedCacheClusterTopicsResponseWithDefaults instantiates a new MsgVpnDistributedCacheClusterTopicsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MsgVpnDistributedCacheClusterTopicsResponse) GetData() []MsgVpnDistributedCacheClusterTopic`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnDistributedCacheClusterTopicsResponse) GetDataOk() (*[]MsgVpnDistributedCacheClusterTopic, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnDistributedCacheClusterTopicsResponse) SetData(v []MsgVpnDistributedCacheClusterTopic)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnDistributedCacheClusterTopicsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnDistributedCacheClusterTopicsResponse) GetLinks() []MsgVpnDistributedCacheClusterTopicLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnDistributedCacheClusterTopicsResponse) GetLinksOk() (*[]MsgVpnDistributedCacheClusterTopicLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnDistributedCacheClusterTopicsResponse) SetLinks(v []MsgVpnDistributedCacheClusterTopicLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnDistributedCacheClusterTopicsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnDistributedCacheClusterTopicsResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnDistributedCacheClusterTopicsResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnDistributedCacheClusterTopicsResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


