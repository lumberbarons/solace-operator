# MsgVpnDistributedCacheClusterInstanceCounter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MsgCount** | Pointer to **int64** | The number of messages cached for the Cache Instance. Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object. | [optional] 
**MsgPeakCount** | Pointer to **int64** | The number of messages cached peak for the Cache Instance. Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object. | [optional] 
**RequestQueueDepthCount** | Pointer to **int64** | The received request message queue depth for the Cache Instance. Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object. | [optional] 
**RequestQueueDepthPeakCount** | Pointer to **int64** | The received request message queue depth peak for the Cache Instance. Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object. | [optional] 
**TopicCount** | Pointer to **int64** | The number of topics cached for the Cache Instance. Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object. | [optional] 
**TopicPeakCount** | Pointer to **int64** | The number of topics cached peak for the Cache Instance. Deprecated since 2.13. This attribute has been moved to the MsgVpnDistributedCacheClusterInstance object. | [optional] 

## Methods

### NewMsgVpnDistributedCacheClusterInstanceCounter

`func NewMsgVpnDistributedCacheClusterInstanceCounter() *MsgVpnDistributedCacheClusterInstanceCounter`

NewMsgVpnDistributedCacheClusterInstanceCounter instantiates a new MsgVpnDistributedCacheClusterInstanceCounter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnDistributedCacheClusterInstanceCounterWithDefaults

`func NewMsgVpnDistributedCacheClusterInstanceCounterWithDefaults() *MsgVpnDistributedCacheClusterInstanceCounter`

NewMsgVpnDistributedCacheClusterInstanceCounterWithDefaults instantiates a new MsgVpnDistributedCacheClusterInstanceCounter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsgCount

`func (o *MsgVpnDistributedCacheClusterInstanceCounter) GetMsgCount() int64`

GetMsgCount returns the MsgCount field if non-nil, zero value otherwise.

### GetMsgCountOk

`func (o *MsgVpnDistributedCacheClusterInstanceCounter) GetMsgCountOk() (*int64, bool)`

GetMsgCountOk returns a tuple with the MsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgCount

`func (o *MsgVpnDistributedCacheClusterInstanceCounter) SetMsgCount(v int64)`

SetMsgCount sets MsgCount field to given value.

### HasMsgCount

`func (o *MsgVpnDistributedCacheClusterInstanceCounter) HasMsgCount() bool`

HasMsgCount returns a boolean if a field has been set.

### GetMsgPeakCount

`func (o *MsgVpnDistributedCacheClusterInstanceCounter) GetMsgPeakCount() int64`

GetMsgPeakCount returns the MsgPeakCount field if non-nil, zero value otherwise.

### GetMsgPeakCountOk

`func (o *MsgVpnDistributedCacheClusterInstanceCounter) GetMsgPeakCountOk() (*int64, bool)`

GetMsgPeakCountOk returns a tuple with the MsgPeakCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgPeakCount

`func (o *MsgVpnDistributedCacheClusterInstanceCounter) SetMsgPeakCount(v int64)`

SetMsgPeakCount sets MsgPeakCount field to given value.

### HasMsgPeakCount

`func (o *MsgVpnDistributedCacheClusterInstanceCounter) HasMsgPeakCount() bool`

HasMsgPeakCount returns a boolean if a field has been set.

### GetRequestQueueDepthCount

`func (o *MsgVpnDistributedCacheClusterInstanceCounter) GetRequestQueueDepthCount() int64`

GetRequestQueueDepthCount returns the RequestQueueDepthCount field if non-nil, zero value otherwise.

### GetRequestQueueDepthCountOk

`func (o *MsgVpnDistributedCacheClusterInstanceCounter) GetRequestQueueDepthCountOk() (*int64, bool)`

GetRequestQueueDepthCountOk returns a tuple with the RequestQueueDepthCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestQueueDepthCount

`func (o *MsgVpnDistributedCacheClusterInstanceCounter) SetRequestQueueDepthCount(v int64)`

SetRequestQueueDepthCount sets RequestQueueDepthCount field to given value.

### HasRequestQueueDepthCount

`func (o *MsgVpnDistributedCacheClusterInstanceCounter) HasRequestQueueDepthCount() bool`

HasRequestQueueDepthCount returns a boolean if a field has been set.

### GetRequestQueueDepthPeakCount

`func (o *MsgVpnDistributedCacheClusterInstanceCounter) GetRequestQueueDepthPeakCount() int64`

GetRequestQueueDepthPeakCount returns the RequestQueueDepthPeakCount field if non-nil, zero value otherwise.

### GetRequestQueueDepthPeakCountOk

`func (o *MsgVpnDistributedCacheClusterInstanceCounter) GetRequestQueueDepthPeakCountOk() (*int64, bool)`

GetRequestQueueDepthPeakCountOk returns a tuple with the RequestQueueDepthPeakCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestQueueDepthPeakCount

`func (o *MsgVpnDistributedCacheClusterInstanceCounter) SetRequestQueueDepthPeakCount(v int64)`

SetRequestQueueDepthPeakCount sets RequestQueueDepthPeakCount field to given value.

### HasRequestQueueDepthPeakCount

`func (o *MsgVpnDistributedCacheClusterInstanceCounter) HasRequestQueueDepthPeakCount() bool`

HasRequestQueueDepthPeakCount returns a boolean if a field has been set.

### GetTopicCount

`func (o *MsgVpnDistributedCacheClusterInstanceCounter) GetTopicCount() int64`

GetTopicCount returns the TopicCount field if non-nil, zero value otherwise.

### GetTopicCountOk

`func (o *MsgVpnDistributedCacheClusterInstanceCounter) GetTopicCountOk() (*int64, bool)`

GetTopicCountOk returns a tuple with the TopicCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicCount

`func (o *MsgVpnDistributedCacheClusterInstanceCounter) SetTopicCount(v int64)`

SetTopicCount sets TopicCount field to given value.

### HasTopicCount

`func (o *MsgVpnDistributedCacheClusterInstanceCounter) HasTopicCount() bool`

HasTopicCount returns a boolean if a field has been set.

### GetTopicPeakCount

`func (o *MsgVpnDistributedCacheClusterInstanceCounter) GetTopicPeakCount() int64`

GetTopicPeakCount returns the TopicPeakCount field if non-nil, zero value otherwise.

### GetTopicPeakCountOk

`func (o *MsgVpnDistributedCacheClusterInstanceCounter) GetTopicPeakCountOk() (*int64, bool)`

GetTopicPeakCountOk returns a tuple with the TopicPeakCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicPeakCount

`func (o *MsgVpnDistributedCacheClusterInstanceCounter) SetTopicPeakCount(v int64)`

SetTopicPeakCount sets TopicPeakCount field to given value.

### HasTopicPeakCount

`func (o *MsgVpnDistributedCacheClusterInstanceCounter) HasTopicPeakCount() bool`

HasTopicPeakCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


