# MsgVpnMqttRetainCache

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupCacheInstance** | Pointer to **string** | The name of the backup Cache Instance associated with this MQTT Retain Cache. | [optional] 
**BackupFailureReason** | Pointer to **string** | The reason why the backup cache associated with this MQTT Retain Cache is operationally down, if any. | [optional] 
**BackupUp** | Pointer to **bool** | Indicates whether the backup cache associated with this MQTT Retain Cache is operationally up. | [optional] 
**BackupUptime** | Pointer to **int32** | The number of seconds that the backup cache associated with this MQTT Retain Cache has been operationally up. | [optional] 
**CacheCluster** | Pointer to **string** | The name of the Cache Cluster associated with this MQTT Retain Cache. | [optional] 
**CacheName** | Pointer to **string** | The name of the MQTT Retain Cache. | [optional] 
**DistributedCache** | Pointer to **string** | The name of the Distributed Cache associated with this MQTT Retain Cache. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether this MQTT Retain Cache is enabled. When the cache is disabled, neither retain messages nor retain requests will be delivered by the cache. However, live retain messages will continue to be delivered to currently connected MQTT clients. | [optional] 
**FailureReason** | Pointer to **string** | The reason why this MQTT Retain Cache is operationally down, if any. | [optional] 
**MsgLifetime** | Pointer to **int64** | The message lifetime, in seconds. If a message remains cached for the duration of its lifetime, the cache will remove the message. A lifetime of 0 results in the message being retained indefinitely. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**PrimaryCacheInstance** | Pointer to **string** | The name of the primary Cache Instance associated with this MQTT Retain Cache. | [optional] 
**PrimaryFailureReason** | Pointer to **string** | The reason why the primary cache associated with this MQTT Retain Cache is operationally down, if any. | [optional] 
**PrimaryUp** | Pointer to **bool** | Indicates whether the primary cache associated with this MQTT Retain Cache is operationally up. | [optional] 
**PrimaryUptime** | Pointer to **int32** | The number of seconds that the primary cache associated with this MQTT Retain Cache has been operationally up. | [optional] 
**Up** | Pointer to **bool** | Indicates whether this MQTT Retain Cache is operationally up. | [optional] 
**Uptime** | Pointer to **int32** | The number of seconds that the MQTT Retain Cache has been operationally up. | [optional] 

## Methods

### NewMsgVpnMqttRetainCache

`func NewMsgVpnMqttRetainCache() *MsgVpnMqttRetainCache`

NewMsgVpnMqttRetainCache instantiates a new MsgVpnMqttRetainCache object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnMqttRetainCacheWithDefaults

`func NewMsgVpnMqttRetainCacheWithDefaults() *MsgVpnMqttRetainCache`

NewMsgVpnMqttRetainCacheWithDefaults instantiates a new MsgVpnMqttRetainCache object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupCacheInstance

`func (o *MsgVpnMqttRetainCache) GetBackupCacheInstance() string`

GetBackupCacheInstance returns the BackupCacheInstance field if non-nil, zero value otherwise.

### GetBackupCacheInstanceOk

`func (o *MsgVpnMqttRetainCache) GetBackupCacheInstanceOk() (*string, bool)`

GetBackupCacheInstanceOk returns a tuple with the BackupCacheInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupCacheInstance

`func (o *MsgVpnMqttRetainCache) SetBackupCacheInstance(v string)`

SetBackupCacheInstance sets BackupCacheInstance field to given value.

### HasBackupCacheInstance

`func (o *MsgVpnMqttRetainCache) HasBackupCacheInstance() bool`

HasBackupCacheInstance returns a boolean if a field has been set.

### GetBackupFailureReason

`func (o *MsgVpnMqttRetainCache) GetBackupFailureReason() string`

GetBackupFailureReason returns the BackupFailureReason field if non-nil, zero value otherwise.

### GetBackupFailureReasonOk

`func (o *MsgVpnMqttRetainCache) GetBackupFailureReasonOk() (*string, bool)`

GetBackupFailureReasonOk returns a tuple with the BackupFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupFailureReason

`func (o *MsgVpnMqttRetainCache) SetBackupFailureReason(v string)`

SetBackupFailureReason sets BackupFailureReason field to given value.

### HasBackupFailureReason

`func (o *MsgVpnMqttRetainCache) HasBackupFailureReason() bool`

HasBackupFailureReason returns a boolean if a field has been set.

### GetBackupUp

`func (o *MsgVpnMqttRetainCache) GetBackupUp() bool`

GetBackupUp returns the BackupUp field if non-nil, zero value otherwise.

### GetBackupUpOk

`func (o *MsgVpnMqttRetainCache) GetBackupUpOk() (*bool, bool)`

GetBackupUpOk returns a tuple with the BackupUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupUp

`func (o *MsgVpnMqttRetainCache) SetBackupUp(v bool)`

SetBackupUp sets BackupUp field to given value.

### HasBackupUp

`func (o *MsgVpnMqttRetainCache) HasBackupUp() bool`

HasBackupUp returns a boolean if a field has been set.

### GetBackupUptime

`func (o *MsgVpnMqttRetainCache) GetBackupUptime() int32`

GetBackupUptime returns the BackupUptime field if non-nil, zero value otherwise.

### GetBackupUptimeOk

`func (o *MsgVpnMqttRetainCache) GetBackupUptimeOk() (*int32, bool)`

GetBackupUptimeOk returns a tuple with the BackupUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupUptime

`func (o *MsgVpnMqttRetainCache) SetBackupUptime(v int32)`

SetBackupUptime sets BackupUptime field to given value.

### HasBackupUptime

`func (o *MsgVpnMqttRetainCache) HasBackupUptime() bool`

HasBackupUptime returns a boolean if a field has been set.

### GetCacheCluster

`func (o *MsgVpnMqttRetainCache) GetCacheCluster() string`

GetCacheCluster returns the CacheCluster field if non-nil, zero value otherwise.

### GetCacheClusterOk

`func (o *MsgVpnMqttRetainCache) GetCacheClusterOk() (*string, bool)`

GetCacheClusterOk returns a tuple with the CacheCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheCluster

`func (o *MsgVpnMqttRetainCache) SetCacheCluster(v string)`

SetCacheCluster sets CacheCluster field to given value.

### HasCacheCluster

`func (o *MsgVpnMqttRetainCache) HasCacheCluster() bool`

HasCacheCluster returns a boolean if a field has been set.

### GetCacheName

`func (o *MsgVpnMqttRetainCache) GetCacheName() string`

GetCacheName returns the CacheName field if non-nil, zero value otherwise.

### GetCacheNameOk

`func (o *MsgVpnMqttRetainCache) GetCacheNameOk() (*string, bool)`

GetCacheNameOk returns a tuple with the CacheName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheName

`func (o *MsgVpnMqttRetainCache) SetCacheName(v string)`

SetCacheName sets CacheName field to given value.

### HasCacheName

`func (o *MsgVpnMqttRetainCache) HasCacheName() bool`

HasCacheName returns a boolean if a field has been set.

### GetDistributedCache

`func (o *MsgVpnMqttRetainCache) GetDistributedCache() string`

GetDistributedCache returns the DistributedCache field if non-nil, zero value otherwise.

### GetDistributedCacheOk

`func (o *MsgVpnMqttRetainCache) GetDistributedCacheOk() (*string, bool)`

GetDistributedCacheOk returns a tuple with the DistributedCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributedCache

`func (o *MsgVpnMqttRetainCache) SetDistributedCache(v string)`

SetDistributedCache sets DistributedCache field to given value.

### HasDistributedCache

`func (o *MsgVpnMqttRetainCache) HasDistributedCache() bool`

HasDistributedCache returns a boolean if a field has been set.

### GetEnabled

`func (o *MsgVpnMqttRetainCache) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MsgVpnMqttRetainCache) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MsgVpnMqttRetainCache) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MsgVpnMqttRetainCache) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFailureReason

`func (o *MsgVpnMqttRetainCache) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *MsgVpnMqttRetainCache) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *MsgVpnMqttRetainCache) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *MsgVpnMqttRetainCache) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetMsgLifetime

`func (o *MsgVpnMqttRetainCache) GetMsgLifetime() int64`

GetMsgLifetime returns the MsgLifetime field if non-nil, zero value otherwise.

### GetMsgLifetimeOk

`func (o *MsgVpnMqttRetainCache) GetMsgLifetimeOk() (*int64, bool)`

GetMsgLifetimeOk returns a tuple with the MsgLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgLifetime

`func (o *MsgVpnMqttRetainCache) SetMsgLifetime(v int64)`

SetMsgLifetime sets MsgLifetime field to given value.

### HasMsgLifetime

`func (o *MsgVpnMqttRetainCache) HasMsgLifetime() bool`

HasMsgLifetime returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnMqttRetainCache) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnMqttRetainCache) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnMqttRetainCache) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnMqttRetainCache) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetPrimaryCacheInstance

`func (o *MsgVpnMqttRetainCache) GetPrimaryCacheInstance() string`

GetPrimaryCacheInstance returns the PrimaryCacheInstance field if non-nil, zero value otherwise.

### GetPrimaryCacheInstanceOk

`func (o *MsgVpnMqttRetainCache) GetPrimaryCacheInstanceOk() (*string, bool)`

GetPrimaryCacheInstanceOk returns a tuple with the PrimaryCacheInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCacheInstance

`func (o *MsgVpnMqttRetainCache) SetPrimaryCacheInstance(v string)`

SetPrimaryCacheInstance sets PrimaryCacheInstance field to given value.

### HasPrimaryCacheInstance

`func (o *MsgVpnMqttRetainCache) HasPrimaryCacheInstance() bool`

HasPrimaryCacheInstance returns a boolean if a field has been set.

### GetPrimaryFailureReason

`func (o *MsgVpnMqttRetainCache) GetPrimaryFailureReason() string`

GetPrimaryFailureReason returns the PrimaryFailureReason field if non-nil, zero value otherwise.

### GetPrimaryFailureReasonOk

`func (o *MsgVpnMqttRetainCache) GetPrimaryFailureReasonOk() (*string, bool)`

GetPrimaryFailureReasonOk returns a tuple with the PrimaryFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryFailureReason

`func (o *MsgVpnMqttRetainCache) SetPrimaryFailureReason(v string)`

SetPrimaryFailureReason sets PrimaryFailureReason field to given value.

### HasPrimaryFailureReason

`func (o *MsgVpnMqttRetainCache) HasPrimaryFailureReason() bool`

HasPrimaryFailureReason returns a boolean if a field has been set.

### GetPrimaryUp

`func (o *MsgVpnMqttRetainCache) GetPrimaryUp() bool`

GetPrimaryUp returns the PrimaryUp field if non-nil, zero value otherwise.

### GetPrimaryUpOk

`func (o *MsgVpnMqttRetainCache) GetPrimaryUpOk() (*bool, bool)`

GetPrimaryUpOk returns a tuple with the PrimaryUp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryUp

`func (o *MsgVpnMqttRetainCache) SetPrimaryUp(v bool)`

SetPrimaryUp sets PrimaryUp field to given value.

### HasPrimaryUp

`func (o *MsgVpnMqttRetainCache) HasPrimaryUp() bool`

HasPrimaryUp returns a boolean if a field has been set.

### GetPrimaryUptime

`func (o *MsgVpnMqttRetainCache) GetPrimaryUptime() int32`

GetPrimaryUptime returns the PrimaryUptime field if non-nil, zero value otherwise.

### GetPrimaryUptimeOk

`func (o *MsgVpnMqttRetainCache) GetPrimaryUptimeOk() (*int32, bool)`

GetPrimaryUptimeOk returns a tuple with the PrimaryUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryUptime

`func (o *MsgVpnMqttRetainCache) SetPrimaryUptime(v int32)`

SetPrimaryUptime sets PrimaryUptime field to given value.

### HasPrimaryUptime

`func (o *MsgVpnMqttRetainCache) HasPrimaryUptime() bool`

HasPrimaryUptime returns a boolean if a field has been set.

### GetUp

`func (o *MsgVpnMqttRetainCache) GetUp() bool`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *MsgVpnMqttRetainCache) GetUpOk() (*bool, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *MsgVpnMqttRetainCache) SetUp(v bool)`

SetUp sets Up field to given value.

### HasUp

`func (o *MsgVpnMqttRetainCache) HasUp() bool`

HasUp returns a boolean if a field has been set.

### GetUptime

`func (o *MsgVpnMqttRetainCache) GetUptime() int32`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *MsgVpnMqttRetainCache) GetUptimeOk() (*int32, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *MsgVpnMqttRetainCache) SetUptime(v int32)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *MsgVpnMqttRetainCache) HasUptime() bool`

HasUptime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


