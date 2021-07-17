# MsgVpnDistributedCacheCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheName** | Pointer to **string** | The name of the Distributed Cache. | [optional] 
**ClusterName** | Pointer to **string** | The name of the Cache Cluster. | [optional] 
**DeliverToOneOverrideEnabled** | Pointer to **bool** | Enable or disable deliver-to-one override for the Cache Cluster. The default value is &#x60;true&#x60;. | [optional] 
**Enabled** | Pointer to **bool** | Enable or disable the Cache Cluster. The default value is &#x60;false&#x60;. | [optional] 
**EventDataByteRateThreshold** | Pointer to [**EventThresholdByValue**](EventThresholdByValue.md) |  | [optional] 
**EventDataMsgRateThreshold** | Pointer to [**EventThresholdByValue**](EventThresholdByValue.md) |  | [optional] 
**EventMaxMemoryThreshold** | Pointer to [**EventThresholdByPercent**](EventThresholdByPercent.md) |  | [optional] 
**EventMaxTopicsThreshold** | Pointer to [**EventThresholdByPercent**](EventThresholdByPercent.md) |  | [optional] 
**EventRequestQueueDepthThreshold** | Pointer to [**EventThresholdByPercent**](EventThresholdByPercent.md) |  | [optional] 
**EventRequestRateThreshold** | Pointer to [**EventThresholdByValue**](EventThresholdByValue.md) |  | [optional] 
**EventResponseRateThreshold** | Pointer to [**EventThresholdByValue**](EventThresholdByValue.md) |  | [optional] 
**GlobalCachingEnabled** | Pointer to **bool** | Enable or disable global caching for the Cache Cluster. When enabled, the Cache Instances will fetch topics from remote Home Cache Clusters when requested, and subscribe to those topics to cache them locally. When disabled, the Cache Instances will remove all subscriptions and cached messages for topics from remote Home Cache Clusters. The default value is &#x60;false&#x60;. | [optional] 
**GlobalCachingHeartbeat** | Pointer to **int64** | The heartbeat interval, in seconds, used by the Cache Instances to monitor connectivity with the remote Home Cache Clusters. The default value is &#x60;3&#x60;. | [optional] 
**GlobalCachingTopicLifetime** | Pointer to **int64** | The topic lifetime, in seconds. If no client requests are received for a given global topic over the duration of the topic lifetime, then the Cache Instance will remove the subscription and cached messages for that topic. A value of 0 disables aging. The default value is &#x60;3600&#x60;. | [optional] 
**MaxMemory** | Pointer to **int64** | The maximum memory usage, in megabytes (MB), for each Cache Instance in the Cache Cluster. The default value is &#x60;2048&#x60;. | [optional] 
**MaxMsgsPerTopic** | Pointer to **int64** | The maximum number of messages per topic for each Cache Instance in the Cache Cluster. When at the maximum, old messages are removed as new messages arrive. The default value is &#x60;1&#x60;. | [optional] 
**MaxRequestQueueDepth** | Pointer to **int64** | The maximum queue depth for cache requests received by the Cache Cluster. The default value is &#x60;100000&#x60;. | [optional] 
**MaxTopicCount** | Pointer to **int64** | The maximum number of topics for each Cache Instance in the Cache Cluster. The default value is &#x60;2000000&#x60;. | [optional] 
**MsgLifetime** | Pointer to **int64** | The message lifetime, in seconds. If a message remains cached for the duration of its lifetime, the Cache Instance will remove the message. A lifetime of 0 results in the message being retained indefinitely. The default value is &#x60;0&#x60;. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**NewTopicAdvertisementEnabled** | Pointer to **bool** | Enable or disable the advertising, onto the message bus, of new topics learned by each Cache Instance in the Cache Cluster. The default value is &#x60;false&#x60;. | [optional] 

## Methods

### NewMsgVpnDistributedCacheCluster

`func NewMsgVpnDistributedCacheCluster() *MsgVpnDistributedCacheCluster`

NewMsgVpnDistributedCacheCluster instantiates a new MsgVpnDistributedCacheCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnDistributedCacheClusterWithDefaults

`func NewMsgVpnDistributedCacheClusterWithDefaults() *MsgVpnDistributedCacheCluster`

NewMsgVpnDistributedCacheClusterWithDefaults instantiates a new MsgVpnDistributedCacheCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheName

`func (o *MsgVpnDistributedCacheCluster) GetCacheName() string`

GetCacheName returns the CacheName field if non-nil, zero value otherwise.

### GetCacheNameOk

`func (o *MsgVpnDistributedCacheCluster) GetCacheNameOk() (*string, bool)`

GetCacheNameOk returns a tuple with the CacheName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheName

`func (o *MsgVpnDistributedCacheCluster) SetCacheName(v string)`

SetCacheName sets CacheName field to given value.

### HasCacheName

`func (o *MsgVpnDistributedCacheCluster) HasCacheName() bool`

HasCacheName returns a boolean if a field has been set.

### GetClusterName

`func (o *MsgVpnDistributedCacheCluster) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *MsgVpnDistributedCacheCluster) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *MsgVpnDistributedCacheCluster) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *MsgVpnDistributedCacheCluster) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetDeliverToOneOverrideEnabled

`func (o *MsgVpnDistributedCacheCluster) GetDeliverToOneOverrideEnabled() bool`

GetDeliverToOneOverrideEnabled returns the DeliverToOneOverrideEnabled field if non-nil, zero value otherwise.

### GetDeliverToOneOverrideEnabledOk

`func (o *MsgVpnDistributedCacheCluster) GetDeliverToOneOverrideEnabledOk() (*bool, bool)`

GetDeliverToOneOverrideEnabledOk returns a tuple with the DeliverToOneOverrideEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverToOneOverrideEnabled

`func (o *MsgVpnDistributedCacheCluster) SetDeliverToOneOverrideEnabled(v bool)`

SetDeliverToOneOverrideEnabled sets DeliverToOneOverrideEnabled field to given value.

### HasDeliverToOneOverrideEnabled

`func (o *MsgVpnDistributedCacheCluster) HasDeliverToOneOverrideEnabled() bool`

HasDeliverToOneOverrideEnabled returns a boolean if a field has been set.

### GetEnabled

`func (o *MsgVpnDistributedCacheCluster) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MsgVpnDistributedCacheCluster) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MsgVpnDistributedCacheCluster) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MsgVpnDistributedCacheCluster) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEventDataByteRateThreshold

`func (o *MsgVpnDistributedCacheCluster) GetEventDataByteRateThreshold() EventThresholdByValue`

GetEventDataByteRateThreshold returns the EventDataByteRateThreshold field if non-nil, zero value otherwise.

### GetEventDataByteRateThresholdOk

`func (o *MsgVpnDistributedCacheCluster) GetEventDataByteRateThresholdOk() (*EventThresholdByValue, bool)`

GetEventDataByteRateThresholdOk returns a tuple with the EventDataByteRateThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDataByteRateThreshold

`func (o *MsgVpnDistributedCacheCluster) SetEventDataByteRateThreshold(v EventThresholdByValue)`

SetEventDataByteRateThreshold sets EventDataByteRateThreshold field to given value.

### HasEventDataByteRateThreshold

`func (o *MsgVpnDistributedCacheCluster) HasEventDataByteRateThreshold() bool`

HasEventDataByteRateThreshold returns a boolean if a field has been set.

### GetEventDataMsgRateThreshold

`func (o *MsgVpnDistributedCacheCluster) GetEventDataMsgRateThreshold() EventThresholdByValue`

GetEventDataMsgRateThreshold returns the EventDataMsgRateThreshold field if non-nil, zero value otherwise.

### GetEventDataMsgRateThresholdOk

`func (o *MsgVpnDistributedCacheCluster) GetEventDataMsgRateThresholdOk() (*EventThresholdByValue, bool)`

GetEventDataMsgRateThresholdOk returns a tuple with the EventDataMsgRateThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDataMsgRateThreshold

`func (o *MsgVpnDistributedCacheCluster) SetEventDataMsgRateThreshold(v EventThresholdByValue)`

SetEventDataMsgRateThreshold sets EventDataMsgRateThreshold field to given value.

### HasEventDataMsgRateThreshold

`func (o *MsgVpnDistributedCacheCluster) HasEventDataMsgRateThreshold() bool`

HasEventDataMsgRateThreshold returns a boolean if a field has been set.

### GetEventMaxMemoryThreshold

`func (o *MsgVpnDistributedCacheCluster) GetEventMaxMemoryThreshold() EventThresholdByPercent`

GetEventMaxMemoryThreshold returns the EventMaxMemoryThreshold field if non-nil, zero value otherwise.

### GetEventMaxMemoryThresholdOk

`func (o *MsgVpnDistributedCacheCluster) GetEventMaxMemoryThresholdOk() (*EventThresholdByPercent, bool)`

GetEventMaxMemoryThresholdOk returns a tuple with the EventMaxMemoryThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventMaxMemoryThreshold

`func (o *MsgVpnDistributedCacheCluster) SetEventMaxMemoryThreshold(v EventThresholdByPercent)`

SetEventMaxMemoryThreshold sets EventMaxMemoryThreshold field to given value.

### HasEventMaxMemoryThreshold

`func (o *MsgVpnDistributedCacheCluster) HasEventMaxMemoryThreshold() bool`

HasEventMaxMemoryThreshold returns a boolean if a field has been set.

### GetEventMaxTopicsThreshold

`func (o *MsgVpnDistributedCacheCluster) GetEventMaxTopicsThreshold() EventThresholdByPercent`

GetEventMaxTopicsThreshold returns the EventMaxTopicsThreshold field if non-nil, zero value otherwise.

### GetEventMaxTopicsThresholdOk

`func (o *MsgVpnDistributedCacheCluster) GetEventMaxTopicsThresholdOk() (*EventThresholdByPercent, bool)`

GetEventMaxTopicsThresholdOk returns a tuple with the EventMaxTopicsThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventMaxTopicsThreshold

`func (o *MsgVpnDistributedCacheCluster) SetEventMaxTopicsThreshold(v EventThresholdByPercent)`

SetEventMaxTopicsThreshold sets EventMaxTopicsThreshold field to given value.

### HasEventMaxTopicsThreshold

`func (o *MsgVpnDistributedCacheCluster) HasEventMaxTopicsThreshold() bool`

HasEventMaxTopicsThreshold returns a boolean if a field has been set.

### GetEventRequestQueueDepthThreshold

`func (o *MsgVpnDistributedCacheCluster) GetEventRequestQueueDepthThreshold() EventThresholdByPercent`

GetEventRequestQueueDepthThreshold returns the EventRequestQueueDepthThreshold field if non-nil, zero value otherwise.

### GetEventRequestQueueDepthThresholdOk

`func (o *MsgVpnDistributedCacheCluster) GetEventRequestQueueDepthThresholdOk() (*EventThresholdByPercent, bool)`

GetEventRequestQueueDepthThresholdOk returns a tuple with the EventRequestQueueDepthThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventRequestQueueDepthThreshold

`func (o *MsgVpnDistributedCacheCluster) SetEventRequestQueueDepthThreshold(v EventThresholdByPercent)`

SetEventRequestQueueDepthThreshold sets EventRequestQueueDepthThreshold field to given value.

### HasEventRequestQueueDepthThreshold

`func (o *MsgVpnDistributedCacheCluster) HasEventRequestQueueDepthThreshold() bool`

HasEventRequestQueueDepthThreshold returns a boolean if a field has been set.

### GetEventRequestRateThreshold

`func (o *MsgVpnDistributedCacheCluster) GetEventRequestRateThreshold() EventThresholdByValue`

GetEventRequestRateThreshold returns the EventRequestRateThreshold field if non-nil, zero value otherwise.

### GetEventRequestRateThresholdOk

`func (o *MsgVpnDistributedCacheCluster) GetEventRequestRateThresholdOk() (*EventThresholdByValue, bool)`

GetEventRequestRateThresholdOk returns a tuple with the EventRequestRateThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventRequestRateThreshold

`func (o *MsgVpnDistributedCacheCluster) SetEventRequestRateThreshold(v EventThresholdByValue)`

SetEventRequestRateThreshold sets EventRequestRateThreshold field to given value.

### HasEventRequestRateThreshold

`func (o *MsgVpnDistributedCacheCluster) HasEventRequestRateThreshold() bool`

HasEventRequestRateThreshold returns a boolean if a field has been set.

### GetEventResponseRateThreshold

`func (o *MsgVpnDistributedCacheCluster) GetEventResponseRateThreshold() EventThresholdByValue`

GetEventResponseRateThreshold returns the EventResponseRateThreshold field if non-nil, zero value otherwise.

### GetEventResponseRateThresholdOk

`func (o *MsgVpnDistributedCacheCluster) GetEventResponseRateThresholdOk() (*EventThresholdByValue, bool)`

GetEventResponseRateThresholdOk returns a tuple with the EventResponseRateThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventResponseRateThreshold

`func (o *MsgVpnDistributedCacheCluster) SetEventResponseRateThreshold(v EventThresholdByValue)`

SetEventResponseRateThreshold sets EventResponseRateThreshold field to given value.

### HasEventResponseRateThreshold

`func (o *MsgVpnDistributedCacheCluster) HasEventResponseRateThreshold() bool`

HasEventResponseRateThreshold returns a boolean if a field has been set.

### GetGlobalCachingEnabled

`func (o *MsgVpnDistributedCacheCluster) GetGlobalCachingEnabled() bool`

GetGlobalCachingEnabled returns the GlobalCachingEnabled field if non-nil, zero value otherwise.

### GetGlobalCachingEnabledOk

`func (o *MsgVpnDistributedCacheCluster) GetGlobalCachingEnabledOk() (*bool, bool)`

GetGlobalCachingEnabledOk returns a tuple with the GlobalCachingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalCachingEnabled

`func (o *MsgVpnDistributedCacheCluster) SetGlobalCachingEnabled(v bool)`

SetGlobalCachingEnabled sets GlobalCachingEnabled field to given value.

### HasGlobalCachingEnabled

`func (o *MsgVpnDistributedCacheCluster) HasGlobalCachingEnabled() bool`

HasGlobalCachingEnabled returns a boolean if a field has been set.

### GetGlobalCachingHeartbeat

`func (o *MsgVpnDistributedCacheCluster) GetGlobalCachingHeartbeat() int64`

GetGlobalCachingHeartbeat returns the GlobalCachingHeartbeat field if non-nil, zero value otherwise.

### GetGlobalCachingHeartbeatOk

`func (o *MsgVpnDistributedCacheCluster) GetGlobalCachingHeartbeatOk() (*int64, bool)`

GetGlobalCachingHeartbeatOk returns a tuple with the GlobalCachingHeartbeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalCachingHeartbeat

`func (o *MsgVpnDistributedCacheCluster) SetGlobalCachingHeartbeat(v int64)`

SetGlobalCachingHeartbeat sets GlobalCachingHeartbeat field to given value.

### HasGlobalCachingHeartbeat

`func (o *MsgVpnDistributedCacheCluster) HasGlobalCachingHeartbeat() bool`

HasGlobalCachingHeartbeat returns a boolean if a field has been set.

### GetGlobalCachingTopicLifetime

`func (o *MsgVpnDistributedCacheCluster) GetGlobalCachingTopicLifetime() int64`

GetGlobalCachingTopicLifetime returns the GlobalCachingTopicLifetime field if non-nil, zero value otherwise.

### GetGlobalCachingTopicLifetimeOk

`func (o *MsgVpnDistributedCacheCluster) GetGlobalCachingTopicLifetimeOk() (*int64, bool)`

GetGlobalCachingTopicLifetimeOk returns a tuple with the GlobalCachingTopicLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalCachingTopicLifetime

`func (o *MsgVpnDistributedCacheCluster) SetGlobalCachingTopicLifetime(v int64)`

SetGlobalCachingTopicLifetime sets GlobalCachingTopicLifetime field to given value.

### HasGlobalCachingTopicLifetime

`func (o *MsgVpnDistributedCacheCluster) HasGlobalCachingTopicLifetime() bool`

HasGlobalCachingTopicLifetime returns a boolean if a field has been set.

### GetMaxMemory

`func (o *MsgVpnDistributedCacheCluster) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *MsgVpnDistributedCacheCluster) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *MsgVpnDistributedCacheCluster) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *MsgVpnDistributedCacheCluster) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxMsgsPerTopic

`func (o *MsgVpnDistributedCacheCluster) GetMaxMsgsPerTopic() int64`

GetMaxMsgsPerTopic returns the MaxMsgsPerTopic field if non-nil, zero value otherwise.

### GetMaxMsgsPerTopicOk

`func (o *MsgVpnDistributedCacheCluster) GetMaxMsgsPerTopicOk() (*int64, bool)`

GetMaxMsgsPerTopicOk returns a tuple with the MaxMsgsPerTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMsgsPerTopic

`func (o *MsgVpnDistributedCacheCluster) SetMaxMsgsPerTopic(v int64)`

SetMaxMsgsPerTopic sets MaxMsgsPerTopic field to given value.

### HasMaxMsgsPerTopic

`func (o *MsgVpnDistributedCacheCluster) HasMaxMsgsPerTopic() bool`

HasMaxMsgsPerTopic returns a boolean if a field has been set.

### GetMaxRequestQueueDepth

`func (o *MsgVpnDistributedCacheCluster) GetMaxRequestQueueDepth() int64`

GetMaxRequestQueueDepth returns the MaxRequestQueueDepth field if non-nil, zero value otherwise.

### GetMaxRequestQueueDepthOk

`func (o *MsgVpnDistributedCacheCluster) GetMaxRequestQueueDepthOk() (*int64, bool)`

GetMaxRequestQueueDepthOk returns a tuple with the MaxRequestQueueDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRequestQueueDepth

`func (o *MsgVpnDistributedCacheCluster) SetMaxRequestQueueDepth(v int64)`

SetMaxRequestQueueDepth sets MaxRequestQueueDepth field to given value.

### HasMaxRequestQueueDepth

`func (o *MsgVpnDistributedCacheCluster) HasMaxRequestQueueDepth() bool`

HasMaxRequestQueueDepth returns a boolean if a field has been set.

### GetMaxTopicCount

`func (o *MsgVpnDistributedCacheCluster) GetMaxTopicCount() int64`

GetMaxTopicCount returns the MaxTopicCount field if non-nil, zero value otherwise.

### GetMaxTopicCountOk

`func (o *MsgVpnDistributedCacheCluster) GetMaxTopicCountOk() (*int64, bool)`

GetMaxTopicCountOk returns a tuple with the MaxTopicCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTopicCount

`func (o *MsgVpnDistributedCacheCluster) SetMaxTopicCount(v int64)`

SetMaxTopicCount sets MaxTopicCount field to given value.

### HasMaxTopicCount

`func (o *MsgVpnDistributedCacheCluster) HasMaxTopicCount() bool`

HasMaxTopicCount returns a boolean if a field has been set.

### GetMsgLifetime

`func (o *MsgVpnDistributedCacheCluster) GetMsgLifetime() int64`

GetMsgLifetime returns the MsgLifetime field if non-nil, zero value otherwise.

### GetMsgLifetimeOk

`func (o *MsgVpnDistributedCacheCluster) GetMsgLifetimeOk() (*int64, bool)`

GetMsgLifetimeOk returns a tuple with the MsgLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgLifetime

`func (o *MsgVpnDistributedCacheCluster) SetMsgLifetime(v int64)`

SetMsgLifetime sets MsgLifetime field to given value.

### HasMsgLifetime

`func (o *MsgVpnDistributedCacheCluster) HasMsgLifetime() bool`

HasMsgLifetime returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnDistributedCacheCluster) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnDistributedCacheCluster) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnDistributedCacheCluster) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnDistributedCacheCluster) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetNewTopicAdvertisementEnabled

`func (o *MsgVpnDistributedCacheCluster) GetNewTopicAdvertisementEnabled() bool`

GetNewTopicAdvertisementEnabled returns the NewTopicAdvertisementEnabled field if non-nil, zero value otherwise.

### GetNewTopicAdvertisementEnabledOk

`func (o *MsgVpnDistributedCacheCluster) GetNewTopicAdvertisementEnabledOk() (*bool, bool)`

GetNewTopicAdvertisementEnabledOk returns a tuple with the NewTopicAdvertisementEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTopicAdvertisementEnabled

`func (o *MsgVpnDistributedCacheCluster) SetNewTopicAdvertisementEnabled(v bool)`

SetNewTopicAdvertisementEnabled sets NewTopicAdvertisementEnabled field to given value.

### HasNewTopicAdvertisementEnabled

`func (o *MsgVpnDistributedCacheCluster) HasNewTopicAdvertisementEnabled() bool`

HasNewTopicAdvertisementEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


