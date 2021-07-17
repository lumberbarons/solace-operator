# MsgVpnDistributedCacheClusterInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoStartEnabled** | Pointer to **bool** | Indicates whether auto-start for the Cache Instance is enabled, and the Cache Instance will automatically attempt to transition from the Stopped operational state to Up whenever it restarts or reconnects to the message broker. | [optional] 
**AverageDataRxBytePeakRate** | Pointer to **int64** | The peak of the one minute average of the data message rate received by the Cache Instance, in bytes per second (B/sec). Available since 2.13. | [optional] 
**AverageDataRxByteRate** | Pointer to **int64** | The one minute average of the data message rate received by the Cache Instance, in bytes per second (B/sec). Available since 2.13. | [optional] 
**AverageDataRxMsgPeakRate** | Pointer to **int64** | The peak of the one minute average of the data message rate received by the Cache Instance, in messages per second (msg/sec). Available since 2.13. | [optional] 
**AverageDataRxMsgRate** | Pointer to **int64** | The one minute average of the data message rate received by the Cache Instance, in messages per second (msg/sec). Available since 2.13. | [optional] 
**AverageDataTxMsgPeakRate** | Pointer to **int64** | The peak of the one minute average of the data message rate transmitted by the Cache Instance, in messages per second (msg/sec). Available since 2.13. | [optional] 
**AverageDataTxMsgRate** | Pointer to **int64** | The one minute average of the data message rate transmitted by the Cache Instance, in messages per second (msg/sec). Available since 2.13. | [optional] 
**AverageRequestRxPeakRate** | Pointer to **int64** | The peak of the one minute average of the request rate received by the Cache Instance, in requests per second (req/sec). Available since 2.13. | [optional] 
**AverageRequestRxRate** | Pointer to **int64** | The one minute average of the request rate received by the Cache Instance, in requests per second (req/sec). Available since 2.13. | [optional] 
**CacheName** | Pointer to **string** | The name of the Distributed Cache. | [optional] 
**ClusterName** | Pointer to **string** | The name of the Cache Cluster. | [optional] 
**Counter** | Pointer to [**MsgVpnDistributedCacheClusterInstanceCounter**](MsgVpnDistributedCacheClusterInstanceCounter.md) |  | [optional] 
**DataRxBytePeakRate** | Pointer to **int64** | The data message peak rate received by the Cache Instance, in bytes per second (B/sec). Available since 2.13. | [optional] 
**DataRxByteRate** | Pointer to **int64** | The data message rate received by the Cache Instance, in bytes per second (B/sec). Available since 2.13. | [optional] 
**DataRxMsgPeakRate** | Pointer to **int64** | The data message peak rate received by the Cache Instance, in messages per second (msg/sec). Available since 2.13. | [optional] 
**DataRxMsgRate** | Pointer to **int64** | The data message rate received by the Cache Instance, in messages per second (msg/sec). Available since 2.13. | [optional] 
**DataTxMsgPeakRate** | Pointer to **int64** | The data message peak rate transmitted by the Cache Instance, in messages per second (msg/sec). Available since 2.13. | [optional] 
**DataTxMsgRate** | Pointer to **int64** | The data message rate transmitted by the Cache Instance, in messages per second (msg/sec). Available since 2.13. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the Cache Instance is enabled. | [optional] 
**InstanceName** | Pointer to **string** | The name of the Cache Instance. | [optional] 
**LastRegisteredTime** | Pointer to **int32** | The timestamp of when the Cache Instance last registered with the message broker. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] 
**LastRxHeartbeatTime** | Pointer to **int32** | The timestamp of the last heartbeat message received from the Cache Instance. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] 
**LastRxSetLostMsgTime** | Pointer to **int32** | The timestamp of the last request for setting the lost message indication received from the Cache Instance. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] 
**LastStoppedReason** | Pointer to **string** | The reason why the Cache Instance was last stopped. | [optional] 
**LastStoppedTime** | Pointer to **int32** | The timestamp of when the Cache Instance was last stopped. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] 
**LastTxClearLostMsgTime** | Pointer to **int32** | The timestamp of the last request for clearing the lost message indication transmitted to the Cache Instance. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] 
**MemoryUsage** | Pointer to **int32** | The memory usage of the Cache Instance, in megabytes (MB). | [optional] 
**MsgCount** | Pointer to **int64** | The number of messages cached for the Cache Instance. Available since 2.13. | [optional] 
**MsgPeakCount** | Pointer to **int64** | The number of messages cached peak for the Cache Instance. Available since 2.13. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**MsgsLost** | Pointer to **bool** | Indicates whether one or more messages were lost by the Cache Instance. | [optional] 
**Rate** | Pointer to [**MsgVpnDistributedCacheClusterInstanceRate**](MsgVpnDistributedCacheClusterInstanceRate.md) |  | [optional] 
**RequestQueueDepthCount** | Pointer to **int64** | The received request message queue depth for the Cache Instance. Available since 2.13. | [optional] 
**RequestQueueDepthPeakCount** | Pointer to **int64** | The received request message queue depth peak for the Cache Instance. Available since 2.13. | [optional] 
**RequestRxPeakRate** | Pointer to **int64** | The request peak rate received by the Cache Instance, in requests per second (req/sec). Available since 2.13. | [optional] 
**RequestRxRate** | Pointer to **int64** | The request rate received by the Cache Instance, in requests per second (req/sec). Available since 2.13. | [optional] 
**State** | Pointer to **string** | The operational state of the Cache Instance. The allowed values and their meaning are:  &lt;pre&gt; \&quot;invalid\&quot; - The Cache Instance state is invalid. \&quot;down\&quot; - The Cache Instance is operationally down. \&quot;stopped\&quot; - The Cache Instance has stopped processing cache requests. \&quot;stopped-lost-msg\&quot; - The Cache Instance has stopped due to a lost message. \&quot;register\&quot; - The Cache Instance is registering with the broker. \&quot;config-sync\&quot; - The Cache Instance is synchronizing its configuration with the broker. \&quot;cluster-sync\&quot; - The Cache Instance is synchronizing its messages with the Cache Cluster. \&quot;up\&quot; - The Cache Instance is operationally up. \&quot;backup\&quot; - The Cache Instance is backing up its messages to disk. \&quot;restore\&quot; - The Cache Instance is restoring its messages from disk. \&quot;not-available\&quot; - The Cache Instance state is not available. &lt;/pre&gt;  | [optional] 
**StopOnLostMsgEnabled** | Pointer to **bool** | Indicates whether stop-on-lost-message is enabled, and the Cache Instance will transition to the Stopped operational state upon losing a message. When Stopped, it cannot accept or respond to cache requests, but continues to cache messages. | [optional] 
**TopicCount** | Pointer to **int64** | The number of topics cached for the Cache Instance. Available since 2.13. | [optional] 
**TopicPeakCount** | Pointer to **int64** | The number of topics cached peak for the Cache Instance. Available since 2.13. | [optional] 

## Methods

### NewMsgVpnDistributedCacheClusterInstance

`func NewMsgVpnDistributedCacheClusterInstance() *MsgVpnDistributedCacheClusterInstance`

NewMsgVpnDistributedCacheClusterInstance instantiates a new MsgVpnDistributedCacheClusterInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnDistributedCacheClusterInstanceWithDefaults

`func NewMsgVpnDistributedCacheClusterInstanceWithDefaults() *MsgVpnDistributedCacheClusterInstance`

NewMsgVpnDistributedCacheClusterInstanceWithDefaults instantiates a new MsgVpnDistributedCacheClusterInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoStartEnabled

`func (o *MsgVpnDistributedCacheClusterInstance) GetAutoStartEnabled() bool`

GetAutoStartEnabled returns the AutoStartEnabled field if non-nil, zero value otherwise.

### GetAutoStartEnabledOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetAutoStartEnabledOk() (*bool, bool)`

GetAutoStartEnabledOk returns a tuple with the AutoStartEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoStartEnabled

`func (o *MsgVpnDistributedCacheClusterInstance) SetAutoStartEnabled(v bool)`

SetAutoStartEnabled sets AutoStartEnabled field to given value.

### HasAutoStartEnabled

`func (o *MsgVpnDistributedCacheClusterInstance) HasAutoStartEnabled() bool`

HasAutoStartEnabled returns a boolean if a field has been set.

### GetAverageDataRxBytePeakRate

`func (o *MsgVpnDistributedCacheClusterInstance) GetAverageDataRxBytePeakRate() int64`

GetAverageDataRxBytePeakRate returns the AverageDataRxBytePeakRate field if non-nil, zero value otherwise.

### GetAverageDataRxBytePeakRateOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetAverageDataRxBytePeakRateOk() (*int64, bool)`

GetAverageDataRxBytePeakRateOk returns a tuple with the AverageDataRxBytePeakRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageDataRxBytePeakRate

`func (o *MsgVpnDistributedCacheClusterInstance) SetAverageDataRxBytePeakRate(v int64)`

SetAverageDataRxBytePeakRate sets AverageDataRxBytePeakRate field to given value.

### HasAverageDataRxBytePeakRate

`func (o *MsgVpnDistributedCacheClusterInstance) HasAverageDataRxBytePeakRate() bool`

HasAverageDataRxBytePeakRate returns a boolean if a field has been set.

### GetAverageDataRxByteRate

`func (o *MsgVpnDistributedCacheClusterInstance) GetAverageDataRxByteRate() int64`

GetAverageDataRxByteRate returns the AverageDataRxByteRate field if non-nil, zero value otherwise.

### GetAverageDataRxByteRateOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetAverageDataRxByteRateOk() (*int64, bool)`

GetAverageDataRxByteRateOk returns a tuple with the AverageDataRxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageDataRxByteRate

`func (o *MsgVpnDistributedCacheClusterInstance) SetAverageDataRxByteRate(v int64)`

SetAverageDataRxByteRate sets AverageDataRxByteRate field to given value.

### HasAverageDataRxByteRate

`func (o *MsgVpnDistributedCacheClusterInstance) HasAverageDataRxByteRate() bool`

HasAverageDataRxByteRate returns a boolean if a field has been set.

### GetAverageDataRxMsgPeakRate

`func (o *MsgVpnDistributedCacheClusterInstance) GetAverageDataRxMsgPeakRate() int64`

GetAverageDataRxMsgPeakRate returns the AverageDataRxMsgPeakRate field if non-nil, zero value otherwise.

### GetAverageDataRxMsgPeakRateOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetAverageDataRxMsgPeakRateOk() (*int64, bool)`

GetAverageDataRxMsgPeakRateOk returns a tuple with the AverageDataRxMsgPeakRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageDataRxMsgPeakRate

`func (o *MsgVpnDistributedCacheClusterInstance) SetAverageDataRxMsgPeakRate(v int64)`

SetAverageDataRxMsgPeakRate sets AverageDataRxMsgPeakRate field to given value.

### HasAverageDataRxMsgPeakRate

`func (o *MsgVpnDistributedCacheClusterInstance) HasAverageDataRxMsgPeakRate() bool`

HasAverageDataRxMsgPeakRate returns a boolean if a field has been set.

### GetAverageDataRxMsgRate

`func (o *MsgVpnDistributedCacheClusterInstance) GetAverageDataRxMsgRate() int64`

GetAverageDataRxMsgRate returns the AverageDataRxMsgRate field if non-nil, zero value otherwise.

### GetAverageDataRxMsgRateOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetAverageDataRxMsgRateOk() (*int64, bool)`

GetAverageDataRxMsgRateOk returns a tuple with the AverageDataRxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageDataRxMsgRate

`func (o *MsgVpnDistributedCacheClusterInstance) SetAverageDataRxMsgRate(v int64)`

SetAverageDataRxMsgRate sets AverageDataRxMsgRate field to given value.

### HasAverageDataRxMsgRate

`func (o *MsgVpnDistributedCacheClusterInstance) HasAverageDataRxMsgRate() bool`

HasAverageDataRxMsgRate returns a boolean if a field has been set.

### GetAverageDataTxMsgPeakRate

`func (o *MsgVpnDistributedCacheClusterInstance) GetAverageDataTxMsgPeakRate() int64`

GetAverageDataTxMsgPeakRate returns the AverageDataTxMsgPeakRate field if non-nil, zero value otherwise.

### GetAverageDataTxMsgPeakRateOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetAverageDataTxMsgPeakRateOk() (*int64, bool)`

GetAverageDataTxMsgPeakRateOk returns a tuple with the AverageDataTxMsgPeakRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageDataTxMsgPeakRate

`func (o *MsgVpnDistributedCacheClusterInstance) SetAverageDataTxMsgPeakRate(v int64)`

SetAverageDataTxMsgPeakRate sets AverageDataTxMsgPeakRate field to given value.

### HasAverageDataTxMsgPeakRate

`func (o *MsgVpnDistributedCacheClusterInstance) HasAverageDataTxMsgPeakRate() bool`

HasAverageDataTxMsgPeakRate returns a boolean if a field has been set.

### GetAverageDataTxMsgRate

`func (o *MsgVpnDistributedCacheClusterInstance) GetAverageDataTxMsgRate() int64`

GetAverageDataTxMsgRate returns the AverageDataTxMsgRate field if non-nil, zero value otherwise.

### GetAverageDataTxMsgRateOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetAverageDataTxMsgRateOk() (*int64, bool)`

GetAverageDataTxMsgRateOk returns a tuple with the AverageDataTxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageDataTxMsgRate

`func (o *MsgVpnDistributedCacheClusterInstance) SetAverageDataTxMsgRate(v int64)`

SetAverageDataTxMsgRate sets AverageDataTxMsgRate field to given value.

### HasAverageDataTxMsgRate

`func (o *MsgVpnDistributedCacheClusterInstance) HasAverageDataTxMsgRate() bool`

HasAverageDataTxMsgRate returns a boolean if a field has been set.

### GetAverageRequestRxPeakRate

`func (o *MsgVpnDistributedCacheClusterInstance) GetAverageRequestRxPeakRate() int64`

GetAverageRequestRxPeakRate returns the AverageRequestRxPeakRate field if non-nil, zero value otherwise.

### GetAverageRequestRxPeakRateOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetAverageRequestRxPeakRateOk() (*int64, bool)`

GetAverageRequestRxPeakRateOk returns a tuple with the AverageRequestRxPeakRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRequestRxPeakRate

`func (o *MsgVpnDistributedCacheClusterInstance) SetAverageRequestRxPeakRate(v int64)`

SetAverageRequestRxPeakRate sets AverageRequestRxPeakRate field to given value.

### HasAverageRequestRxPeakRate

`func (o *MsgVpnDistributedCacheClusterInstance) HasAverageRequestRxPeakRate() bool`

HasAverageRequestRxPeakRate returns a boolean if a field has been set.

### GetAverageRequestRxRate

`func (o *MsgVpnDistributedCacheClusterInstance) GetAverageRequestRxRate() int64`

GetAverageRequestRxRate returns the AverageRequestRxRate field if non-nil, zero value otherwise.

### GetAverageRequestRxRateOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetAverageRequestRxRateOk() (*int64, bool)`

GetAverageRequestRxRateOk returns a tuple with the AverageRequestRxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRequestRxRate

`func (o *MsgVpnDistributedCacheClusterInstance) SetAverageRequestRxRate(v int64)`

SetAverageRequestRxRate sets AverageRequestRxRate field to given value.

### HasAverageRequestRxRate

`func (o *MsgVpnDistributedCacheClusterInstance) HasAverageRequestRxRate() bool`

HasAverageRequestRxRate returns a boolean if a field has been set.

### GetCacheName

`func (o *MsgVpnDistributedCacheClusterInstance) GetCacheName() string`

GetCacheName returns the CacheName field if non-nil, zero value otherwise.

### GetCacheNameOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetCacheNameOk() (*string, bool)`

GetCacheNameOk returns a tuple with the CacheName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheName

`func (o *MsgVpnDistributedCacheClusterInstance) SetCacheName(v string)`

SetCacheName sets CacheName field to given value.

### HasCacheName

`func (o *MsgVpnDistributedCacheClusterInstance) HasCacheName() bool`

HasCacheName returns a boolean if a field has been set.

### GetClusterName

`func (o *MsgVpnDistributedCacheClusterInstance) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *MsgVpnDistributedCacheClusterInstance) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *MsgVpnDistributedCacheClusterInstance) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetCounter

`func (o *MsgVpnDistributedCacheClusterInstance) GetCounter() MsgVpnDistributedCacheClusterInstanceCounter`

GetCounter returns the Counter field if non-nil, zero value otherwise.

### GetCounterOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetCounterOk() (*MsgVpnDistributedCacheClusterInstanceCounter, bool)`

GetCounterOk returns a tuple with the Counter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounter

`func (o *MsgVpnDistributedCacheClusterInstance) SetCounter(v MsgVpnDistributedCacheClusterInstanceCounter)`

SetCounter sets Counter field to given value.

### HasCounter

`func (o *MsgVpnDistributedCacheClusterInstance) HasCounter() bool`

HasCounter returns a boolean if a field has been set.

### GetDataRxBytePeakRate

`func (o *MsgVpnDistributedCacheClusterInstance) GetDataRxBytePeakRate() int64`

GetDataRxBytePeakRate returns the DataRxBytePeakRate field if non-nil, zero value otherwise.

### GetDataRxBytePeakRateOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetDataRxBytePeakRateOk() (*int64, bool)`

GetDataRxBytePeakRateOk returns a tuple with the DataRxBytePeakRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRxBytePeakRate

`func (o *MsgVpnDistributedCacheClusterInstance) SetDataRxBytePeakRate(v int64)`

SetDataRxBytePeakRate sets DataRxBytePeakRate field to given value.

### HasDataRxBytePeakRate

`func (o *MsgVpnDistributedCacheClusterInstance) HasDataRxBytePeakRate() bool`

HasDataRxBytePeakRate returns a boolean if a field has been set.

### GetDataRxByteRate

`func (o *MsgVpnDistributedCacheClusterInstance) GetDataRxByteRate() int64`

GetDataRxByteRate returns the DataRxByteRate field if non-nil, zero value otherwise.

### GetDataRxByteRateOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetDataRxByteRateOk() (*int64, bool)`

GetDataRxByteRateOk returns a tuple with the DataRxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRxByteRate

`func (o *MsgVpnDistributedCacheClusterInstance) SetDataRxByteRate(v int64)`

SetDataRxByteRate sets DataRxByteRate field to given value.

### HasDataRxByteRate

`func (o *MsgVpnDistributedCacheClusterInstance) HasDataRxByteRate() bool`

HasDataRxByteRate returns a boolean if a field has been set.

### GetDataRxMsgPeakRate

`func (o *MsgVpnDistributedCacheClusterInstance) GetDataRxMsgPeakRate() int64`

GetDataRxMsgPeakRate returns the DataRxMsgPeakRate field if non-nil, zero value otherwise.

### GetDataRxMsgPeakRateOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetDataRxMsgPeakRateOk() (*int64, bool)`

GetDataRxMsgPeakRateOk returns a tuple with the DataRxMsgPeakRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRxMsgPeakRate

`func (o *MsgVpnDistributedCacheClusterInstance) SetDataRxMsgPeakRate(v int64)`

SetDataRxMsgPeakRate sets DataRxMsgPeakRate field to given value.

### HasDataRxMsgPeakRate

`func (o *MsgVpnDistributedCacheClusterInstance) HasDataRxMsgPeakRate() bool`

HasDataRxMsgPeakRate returns a boolean if a field has been set.

### GetDataRxMsgRate

`func (o *MsgVpnDistributedCacheClusterInstance) GetDataRxMsgRate() int64`

GetDataRxMsgRate returns the DataRxMsgRate field if non-nil, zero value otherwise.

### GetDataRxMsgRateOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetDataRxMsgRateOk() (*int64, bool)`

GetDataRxMsgRateOk returns a tuple with the DataRxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRxMsgRate

`func (o *MsgVpnDistributedCacheClusterInstance) SetDataRxMsgRate(v int64)`

SetDataRxMsgRate sets DataRxMsgRate field to given value.

### HasDataRxMsgRate

`func (o *MsgVpnDistributedCacheClusterInstance) HasDataRxMsgRate() bool`

HasDataRxMsgRate returns a boolean if a field has been set.

### GetDataTxMsgPeakRate

`func (o *MsgVpnDistributedCacheClusterInstance) GetDataTxMsgPeakRate() int64`

GetDataTxMsgPeakRate returns the DataTxMsgPeakRate field if non-nil, zero value otherwise.

### GetDataTxMsgPeakRateOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetDataTxMsgPeakRateOk() (*int64, bool)`

GetDataTxMsgPeakRateOk returns a tuple with the DataTxMsgPeakRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTxMsgPeakRate

`func (o *MsgVpnDistributedCacheClusterInstance) SetDataTxMsgPeakRate(v int64)`

SetDataTxMsgPeakRate sets DataTxMsgPeakRate field to given value.

### HasDataTxMsgPeakRate

`func (o *MsgVpnDistributedCacheClusterInstance) HasDataTxMsgPeakRate() bool`

HasDataTxMsgPeakRate returns a boolean if a field has been set.

### GetDataTxMsgRate

`func (o *MsgVpnDistributedCacheClusterInstance) GetDataTxMsgRate() int64`

GetDataTxMsgRate returns the DataTxMsgRate field if non-nil, zero value otherwise.

### GetDataTxMsgRateOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetDataTxMsgRateOk() (*int64, bool)`

GetDataTxMsgRateOk returns a tuple with the DataTxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTxMsgRate

`func (o *MsgVpnDistributedCacheClusterInstance) SetDataTxMsgRate(v int64)`

SetDataTxMsgRate sets DataTxMsgRate field to given value.

### HasDataTxMsgRate

`func (o *MsgVpnDistributedCacheClusterInstance) HasDataTxMsgRate() bool`

HasDataTxMsgRate returns a boolean if a field has been set.

### GetEnabled

`func (o *MsgVpnDistributedCacheClusterInstance) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MsgVpnDistributedCacheClusterInstance) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MsgVpnDistributedCacheClusterInstance) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetInstanceName

`func (o *MsgVpnDistributedCacheClusterInstance) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *MsgVpnDistributedCacheClusterInstance) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *MsgVpnDistributedCacheClusterInstance) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### GetLastRegisteredTime

`func (o *MsgVpnDistributedCacheClusterInstance) GetLastRegisteredTime() int32`

GetLastRegisteredTime returns the LastRegisteredTime field if non-nil, zero value otherwise.

### GetLastRegisteredTimeOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetLastRegisteredTimeOk() (*int32, bool)`

GetLastRegisteredTimeOk returns a tuple with the LastRegisteredTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRegisteredTime

`func (o *MsgVpnDistributedCacheClusterInstance) SetLastRegisteredTime(v int32)`

SetLastRegisteredTime sets LastRegisteredTime field to given value.

### HasLastRegisteredTime

`func (o *MsgVpnDistributedCacheClusterInstance) HasLastRegisteredTime() bool`

HasLastRegisteredTime returns a boolean if a field has been set.

### GetLastRxHeartbeatTime

`func (o *MsgVpnDistributedCacheClusterInstance) GetLastRxHeartbeatTime() int32`

GetLastRxHeartbeatTime returns the LastRxHeartbeatTime field if non-nil, zero value otherwise.

### GetLastRxHeartbeatTimeOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetLastRxHeartbeatTimeOk() (*int32, bool)`

GetLastRxHeartbeatTimeOk returns a tuple with the LastRxHeartbeatTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRxHeartbeatTime

`func (o *MsgVpnDistributedCacheClusterInstance) SetLastRxHeartbeatTime(v int32)`

SetLastRxHeartbeatTime sets LastRxHeartbeatTime field to given value.

### HasLastRxHeartbeatTime

`func (o *MsgVpnDistributedCacheClusterInstance) HasLastRxHeartbeatTime() bool`

HasLastRxHeartbeatTime returns a boolean if a field has been set.

### GetLastRxSetLostMsgTime

`func (o *MsgVpnDistributedCacheClusterInstance) GetLastRxSetLostMsgTime() int32`

GetLastRxSetLostMsgTime returns the LastRxSetLostMsgTime field if non-nil, zero value otherwise.

### GetLastRxSetLostMsgTimeOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetLastRxSetLostMsgTimeOk() (*int32, bool)`

GetLastRxSetLostMsgTimeOk returns a tuple with the LastRxSetLostMsgTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRxSetLostMsgTime

`func (o *MsgVpnDistributedCacheClusterInstance) SetLastRxSetLostMsgTime(v int32)`

SetLastRxSetLostMsgTime sets LastRxSetLostMsgTime field to given value.

### HasLastRxSetLostMsgTime

`func (o *MsgVpnDistributedCacheClusterInstance) HasLastRxSetLostMsgTime() bool`

HasLastRxSetLostMsgTime returns a boolean if a field has been set.

### GetLastStoppedReason

`func (o *MsgVpnDistributedCacheClusterInstance) GetLastStoppedReason() string`

GetLastStoppedReason returns the LastStoppedReason field if non-nil, zero value otherwise.

### GetLastStoppedReasonOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetLastStoppedReasonOk() (*string, bool)`

GetLastStoppedReasonOk returns a tuple with the LastStoppedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStoppedReason

`func (o *MsgVpnDistributedCacheClusterInstance) SetLastStoppedReason(v string)`

SetLastStoppedReason sets LastStoppedReason field to given value.

### HasLastStoppedReason

`func (o *MsgVpnDistributedCacheClusterInstance) HasLastStoppedReason() bool`

HasLastStoppedReason returns a boolean if a field has been set.

### GetLastStoppedTime

`func (o *MsgVpnDistributedCacheClusterInstance) GetLastStoppedTime() int32`

GetLastStoppedTime returns the LastStoppedTime field if non-nil, zero value otherwise.

### GetLastStoppedTimeOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetLastStoppedTimeOk() (*int32, bool)`

GetLastStoppedTimeOk returns a tuple with the LastStoppedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStoppedTime

`func (o *MsgVpnDistributedCacheClusterInstance) SetLastStoppedTime(v int32)`

SetLastStoppedTime sets LastStoppedTime field to given value.

### HasLastStoppedTime

`func (o *MsgVpnDistributedCacheClusterInstance) HasLastStoppedTime() bool`

HasLastStoppedTime returns a boolean if a field has been set.

### GetLastTxClearLostMsgTime

`func (o *MsgVpnDistributedCacheClusterInstance) GetLastTxClearLostMsgTime() int32`

GetLastTxClearLostMsgTime returns the LastTxClearLostMsgTime field if non-nil, zero value otherwise.

### GetLastTxClearLostMsgTimeOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetLastTxClearLostMsgTimeOk() (*int32, bool)`

GetLastTxClearLostMsgTimeOk returns a tuple with the LastTxClearLostMsgTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTxClearLostMsgTime

`func (o *MsgVpnDistributedCacheClusterInstance) SetLastTxClearLostMsgTime(v int32)`

SetLastTxClearLostMsgTime sets LastTxClearLostMsgTime field to given value.

### HasLastTxClearLostMsgTime

`func (o *MsgVpnDistributedCacheClusterInstance) HasLastTxClearLostMsgTime() bool`

HasLastTxClearLostMsgTime returns a boolean if a field has been set.

### GetMemoryUsage

`func (o *MsgVpnDistributedCacheClusterInstance) GetMemoryUsage() int32`

GetMemoryUsage returns the MemoryUsage field if non-nil, zero value otherwise.

### GetMemoryUsageOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetMemoryUsageOk() (*int32, bool)`

GetMemoryUsageOk returns a tuple with the MemoryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsage

`func (o *MsgVpnDistributedCacheClusterInstance) SetMemoryUsage(v int32)`

SetMemoryUsage sets MemoryUsage field to given value.

### HasMemoryUsage

`func (o *MsgVpnDistributedCacheClusterInstance) HasMemoryUsage() bool`

HasMemoryUsage returns a boolean if a field has been set.

### GetMsgCount

`func (o *MsgVpnDistributedCacheClusterInstance) GetMsgCount() int64`

GetMsgCount returns the MsgCount field if non-nil, zero value otherwise.

### GetMsgCountOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetMsgCountOk() (*int64, bool)`

GetMsgCountOk returns a tuple with the MsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgCount

`func (o *MsgVpnDistributedCacheClusterInstance) SetMsgCount(v int64)`

SetMsgCount sets MsgCount field to given value.

### HasMsgCount

`func (o *MsgVpnDistributedCacheClusterInstance) HasMsgCount() bool`

HasMsgCount returns a boolean if a field has been set.

### GetMsgPeakCount

`func (o *MsgVpnDistributedCacheClusterInstance) GetMsgPeakCount() int64`

GetMsgPeakCount returns the MsgPeakCount field if non-nil, zero value otherwise.

### GetMsgPeakCountOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetMsgPeakCountOk() (*int64, bool)`

GetMsgPeakCountOk returns a tuple with the MsgPeakCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgPeakCount

`func (o *MsgVpnDistributedCacheClusterInstance) SetMsgPeakCount(v int64)`

SetMsgPeakCount sets MsgPeakCount field to given value.

### HasMsgPeakCount

`func (o *MsgVpnDistributedCacheClusterInstance) HasMsgPeakCount() bool`

HasMsgPeakCount returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnDistributedCacheClusterInstance) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnDistributedCacheClusterInstance) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnDistributedCacheClusterInstance) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetMsgsLost

`func (o *MsgVpnDistributedCacheClusterInstance) GetMsgsLost() bool`

GetMsgsLost returns the MsgsLost field if non-nil, zero value otherwise.

### GetMsgsLostOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetMsgsLostOk() (*bool, bool)`

GetMsgsLostOk returns a tuple with the MsgsLost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgsLost

`func (o *MsgVpnDistributedCacheClusterInstance) SetMsgsLost(v bool)`

SetMsgsLost sets MsgsLost field to given value.

### HasMsgsLost

`func (o *MsgVpnDistributedCacheClusterInstance) HasMsgsLost() bool`

HasMsgsLost returns a boolean if a field has been set.

### GetRate

`func (o *MsgVpnDistributedCacheClusterInstance) GetRate() MsgVpnDistributedCacheClusterInstanceRate`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetRateOk() (*MsgVpnDistributedCacheClusterInstanceRate, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *MsgVpnDistributedCacheClusterInstance) SetRate(v MsgVpnDistributedCacheClusterInstanceRate)`

SetRate sets Rate field to given value.

### HasRate

`func (o *MsgVpnDistributedCacheClusterInstance) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetRequestQueueDepthCount

`func (o *MsgVpnDistributedCacheClusterInstance) GetRequestQueueDepthCount() int64`

GetRequestQueueDepthCount returns the RequestQueueDepthCount field if non-nil, zero value otherwise.

### GetRequestQueueDepthCountOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetRequestQueueDepthCountOk() (*int64, bool)`

GetRequestQueueDepthCountOk returns a tuple with the RequestQueueDepthCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestQueueDepthCount

`func (o *MsgVpnDistributedCacheClusterInstance) SetRequestQueueDepthCount(v int64)`

SetRequestQueueDepthCount sets RequestQueueDepthCount field to given value.

### HasRequestQueueDepthCount

`func (o *MsgVpnDistributedCacheClusterInstance) HasRequestQueueDepthCount() bool`

HasRequestQueueDepthCount returns a boolean if a field has been set.

### GetRequestQueueDepthPeakCount

`func (o *MsgVpnDistributedCacheClusterInstance) GetRequestQueueDepthPeakCount() int64`

GetRequestQueueDepthPeakCount returns the RequestQueueDepthPeakCount field if non-nil, zero value otherwise.

### GetRequestQueueDepthPeakCountOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetRequestQueueDepthPeakCountOk() (*int64, bool)`

GetRequestQueueDepthPeakCountOk returns a tuple with the RequestQueueDepthPeakCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestQueueDepthPeakCount

`func (o *MsgVpnDistributedCacheClusterInstance) SetRequestQueueDepthPeakCount(v int64)`

SetRequestQueueDepthPeakCount sets RequestQueueDepthPeakCount field to given value.

### HasRequestQueueDepthPeakCount

`func (o *MsgVpnDistributedCacheClusterInstance) HasRequestQueueDepthPeakCount() bool`

HasRequestQueueDepthPeakCount returns a boolean if a field has been set.

### GetRequestRxPeakRate

`func (o *MsgVpnDistributedCacheClusterInstance) GetRequestRxPeakRate() int64`

GetRequestRxPeakRate returns the RequestRxPeakRate field if non-nil, zero value otherwise.

### GetRequestRxPeakRateOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetRequestRxPeakRateOk() (*int64, bool)`

GetRequestRxPeakRateOk returns a tuple with the RequestRxPeakRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestRxPeakRate

`func (o *MsgVpnDistributedCacheClusterInstance) SetRequestRxPeakRate(v int64)`

SetRequestRxPeakRate sets RequestRxPeakRate field to given value.

### HasRequestRxPeakRate

`func (o *MsgVpnDistributedCacheClusterInstance) HasRequestRxPeakRate() bool`

HasRequestRxPeakRate returns a boolean if a field has been set.

### GetRequestRxRate

`func (o *MsgVpnDistributedCacheClusterInstance) GetRequestRxRate() int64`

GetRequestRxRate returns the RequestRxRate field if non-nil, zero value otherwise.

### GetRequestRxRateOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetRequestRxRateOk() (*int64, bool)`

GetRequestRxRateOk returns a tuple with the RequestRxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestRxRate

`func (o *MsgVpnDistributedCacheClusterInstance) SetRequestRxRate(v int64)`

SetRequestRxRate sets RequestRxRate field to given value.

### HasRequestRxRate

`func (o *MsgVpnDistributedCacheClusterInstance) HasRequestRxRate() bool`

HasRequestRxRate returns a boolean if a field has been set.

### GetState

`func (o *MsgVpnDistributedCacheClusterInstance) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MsgVpnDistributedCacheClusterInstance) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *MsgVpnDistributedCacheClusterInstance) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStopOnLostMsgEnabled

`func (o *MsgVpnDistributedCacheClusterInstance) GetStopOnLostMsgEnabled() bool`

GetStopOnLostMsgEnabled returns the StopOnLostMsgEnabled field if non-nil, zero value otherwise.

### GetStopOnLostMsgEnabledOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetStopOnLostMsgEnabledOk() (*bool, bool)`

GetStopOnLostMsgEnabledOk returns a tuple with the StopOnLostMsgEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopOnLostMsgEnabled

`func (o *MsgVpnDistributedCacheClusterInstance) SetStopOnLostMsgEnabled(v bool)`

SetStopOnLostMsgEnabled sets StopOnLostMsgEnabled field to given value.

### HasStopOnLostMsgEnabled

`func (o *MsgVpnDistributedCacheClusterInstance) HasStopOnLostMsgEnabled() bool`

HasStopOnLostMsgEnabled returns a boolean if a field has been set.

### GetTopicCount

`func (o *MsgVpnDistributedCacheClusterInstance) GetTopicCount() int64`

GetTopicCount returns the TopicCount field if non-nil, zero value otherwise.

### GetTopicCountOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetTopicCountOk() (*int64, bool)`

GetTopicCountOk returns a tuple with the TopicCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicCount

`func (o *MsgVpnDistributedCacheClusterInstance) SetTopicCount(v int64)`

SetTopicCount sets TopicCount field to given value.

### HasTopicCount

`func (o *MsgVpnDistributedCacheClusterInstance) HasTopicCount() bool`

HasTopicCount returns a boolean if a field has been set.

### GetTopicPeakCount

`func (o *MsgVpnDistributedCacheClusterInstance) GetTopicPeakCount() int64`

GetTopicPeakCount returns the TopicPeakCount field if non-nil, zero value otherwise.

### GetTopicPeakCountOk

`func (o *MsgVpnDistributedCacheClusterInstance) GetTopicPeakCountOk() (*int64, bool)`

GetTopicPeakCountOk returns a tuple with the TopicPeakCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicPeakCount

`func (o *MsgVpnDistributedCacheClusterInstance) SetTopicPeakCount(v int64)`

SetTopicPeakCount sets TopicPeakCount field to given value.

### HasTopicPeakCount

`func (o *MsgVpnDistributedCacheClusterInstance) HasTopicPeakCount() bool`

HasTopicPeakCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


