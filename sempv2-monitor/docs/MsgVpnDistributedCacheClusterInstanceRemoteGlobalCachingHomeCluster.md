# MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheName** | Pointer to **string** | The name of the Distributed Cache. | [optional] 
**CacheRequestForwardedCount** | Pointer to **int64** | The number of cache requests forwarded to the remote Home Cache Cluster. | [optional] 
**CacheRequestRxCount** | Pointer to **int64** | The number of cache requests received from the remote Home Cache Cluster. | [optional] 
**ClusterName** | Pointer to **string** | The name of the Cache Cluster. | [optional] 
**HeartbeatRxCount** | Pointer to **int64** | The number of heartbeat messages received from the remote Home Cache Cluster. | [optional] 
**HeartbeatUp** | Pointer to **bool** | Indicates whether the operational state of the heartbeat with the remote Home Cache Cluster is up. | [optional] 
**HomeClusterName** | Pointer to **string** | The name of the remote Home Cache Cluster. | [optional] 
**InstanceName** | Pointer to **string** | The name of the Cache Instance. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 

## Methods

### NewMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster

`func NewMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster() *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster`

NewMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster instantiates a new MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterWithDefaults

`func NewMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterWithDefaults() *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster`

NewMsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeClusterWithDefaults instantiates a new MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheName

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) GetCacheName() string`

GetCacheName returns the CacheName field if non-nil, zero value otherwise.

### GetCacheNameOk

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) GetCacheNameOk() (*string, bool)`

GetCacheNameOk returns a tuple with the CacheName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheName

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) SetCacheName(v string)`

SetCacheName sets CacheName field to given value.

### HasCacheName

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) HasCacheName() bool`

HasCacheName returns a boolean if a field has been set.

### GetCacheRequestForwardedCount

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) GetCacheRequestForwardedCount() int64`

GetCacheRequestForwardedCount returns the CacheRequestForwardedCount field if non-nil, zero value otherwise.

### GetCacheRequestForwardedCountOk

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) GetCacheRequestForwardedCountOk() (*int64, bool)`

GetCacheRequestForwardedCountOk returns a tuple with the CacheRequestForwardedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheRequestForwardedCount

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) SetCacheRequestForwardedCount(v int64)`

SetCacheRequestForwardedCount sets CacheRequestForwardedCount field to given value.

### HasCacheRequestForwardedCount

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) HasCacheRequestForwardedCount() bool`

HasCacheRequestForwardedCount returns a boolean if a field has been set.

### GetCacheRequestRxCount

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) GetCacheRequestRxCount() int64`

GetCacheRequestRxCount returns the CacheRequestRxCount field if non-nil, zero value otherwise.

### GetCacheRequestRxCountOk

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) GetCacheRequestRxCountOk() (*int64, bool)`

GetCacheRequestRxCountOk returns a tuple with the CacheRequestRxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheRequestRxCount

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) SetCacheRequestRxCount(v int64)`

SetCacheRequestRxCount sets CacheRequestRxCount field to given value.

### HasCacheRequestRxCount

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) HasCacheRequestRxCount() bool`

HasCacheRequestRxCount returns a boolean if a field has been set.

### GetClusterName

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetHeartbeatRxCount

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) GetHeartbeatRxCount() int64`

GetHeartbeatRxCount returns the HeartbeatRxCount field if non-nil, zero value otherwise.

### GetHeartbeatRxCountOk

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) GetHeartbeatRxCountOk() (*int64, bool)`

GetHeartbeatRxCountOk returns a tuple with the HeartbeatRxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartbeatRxCount

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) SetHeartbeatRxCount(v int64)`

SetHeartbeatRxCount sets HeartbeatRxCount field to given value.

### HasHeartbeatRxCount

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) HasHeartbeatRxCount() bool`

HasHeartbeatRxCount returns a boolean if a field has been set.

### GetHeartbeatUp

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) GetHeartbeatUp() bool`

GetHeartbeatUp returns the HeartbeatUp field if non-nil, zero value otherwise.

### GetHeartbeatUpOk

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) GetHeartbeatUpOk() (*bool, bool)`

GetHeartbeatUpOk returns a tuple with the HeartbeatUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartbeatUp

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) SetHeartbeatUp(v bool)`

SetHeartbeatUp sets HeartbeatUp field to given value.

### HasHeartbeatUp

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) HasHeartbeatUp() bool`

HasHeartbeatUp returns a boolean if a field has been set.

### GetHomeClusterName

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) GetHomeClusterName() string`

GetHomeClusterName returns the HomeClusterName field if non-nil, zero value otherwise.

### GetHomeClusterNameOk

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) GetHomeClusterNameOk() (*string, bool)`

GetHomeClusterNameOk returns a tuple with the HomeClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeClusterName

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) SetHomeClusterName(v string)`

SetHomeClusterName sets HomeClusterName field to given value.

### HasHomeClusterName

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) HasHomeClusterName() bool`

HasHomeClusterName returns a boolean if a field has been set.

### GetInstanceName

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteGlobalCachingHomeCluster) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


