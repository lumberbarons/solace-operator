# MsgVpnDistributedCache

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheName** | Pointer to **string** | The name of the Distributed Cache. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the Distributed Cache is enabled. | [optional] 
**Heartbeat** | Pointer to **int64** | The heartbeat interval, in seconds, used by the Cache Instances to monitor connectivity with the message broker. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**MsgsLost** | Pointer to **bool** | Indicates whether one or more messages were lost by any Cache Instance in the Distributed Cache. | [optional] 
**ScheduledDeleteMsgDayList** | Pointer to **string** | The scheduled delete message day(s), specified as \&quot;daily\&quot; or a comma-separated list of days. Days must be specified as \&quot;Sun\&quot;, \&quot;Mon\&quot;, \&quot;Tue\&quot;, \&quot;Wed\&quot;, \&quot;Thu\&quot;, \&quot;Fri\&quot;, or \&quot;Sat\&quot;, with no spaces, and in sorted order from Sunday to Saturday. | [optional] 
**ScheduledDeleteMsgTimeList** | Pointer to **string** | The scheduled delete message time(s), specified as \&quot;hourly\&quot; or a comma-separated list of 24-hour times in the form hh:mm, or h:mm. There must be no spaces, and times must be in sorted order from 0:00 to 23:59. | [optional] 

## Methods

### NewMsgVpnDistributedCache

`func NewMsgVpnDistributedCache() *MsgVpnDistributedCache`

NewMsgVpnDistributedCache instantiates a new MsgVpnDistributedCache object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnDistributedCacheWithDefaults

`func NewMsgVpnDistributedCacheWithDefaults() *MsgVpnDistributedCache`

NewMsgVpnDistributedCacheWithDefaults instantiates a new MsgVpnDistributedCache object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheName

`func (o *MsgVpnDistributedCache) GetCacheName() string`

GetCacheName returns the CacheName field if non-nil, zero value otherwise.

### GetCacheNameOk

`func (o *MsgVpnDistributedCache) GetCacheNameOk() (*string, bool)`

GetCacheNameOk returns a tuple with the CacheName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheName

`func (o *MsgVpnDistributedCache) SetCacheName(v string)`

SetCacheName sets CacheName field to given value.

### HasCacheName

`func (o *MsgVpnDistributedCache) HasCacheName() bool`

HasCacheName returns a boolean if a field has been set.

### GetEnabled

`func (o *MsgVpnDistributedCache) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MsgVpnDistributedCache) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MsgVpnDistributedCache) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MsgVpnDistributedCache) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHeartbeat

`func (o *MsgVpnDistributedCache) GetHeartbeat() int64`

GetHeartbeat returns the Heartbeat field if non-nil, zero value otherwise.

### GetHeartbeatOk

`func (o *MsgVpnDistributedCache) GetHeartbeatOk() (*int64, bool)`

GetHeartbeatOk returns a tuple with the Heartbeat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeartbeat

`func (o *MsgVpnDistributedCache) SetHeartbeat(v int64)`

SetHeartbeat sets Heartbeat field to given value.

### HasHeartbeat

`func (o *MsgVpnDistributedCache) HasHeartbeat() bool`

HasHeartbeat returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnDistributedCache) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnDistributedCache) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnDistributedCache) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnDistributedCache) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetMsgsLost

`func (o *MsgVpnDistributedCache) GetMsgsLost() bool`

GetMsgsLost returns the MsgsLost field if non-nil, zero value otherwise.

### GetMsgsLostOk

`func (o *MsgVpnDistributedCache) GetMsgsLostOk() (*bool, bool)`

GetMsgsLostOk returns a tuple with the MsgsLost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgsLost

`func (o *MsgVpnDistributedCache) SetMsgsLost(v bool)`

SetMsgsLost sets MsgsLost field to given value.

### HasMsgsLost

`func (o *MsgVpnDistributedCache) HasMsgsLost() bool`

HasMsgsLost returns a boolean if a field has been set.

### GetScheduledDeleteMsgDayList

`func (o *MsgVpnDistributedCache) GetScheduledDeleteMsgDayList() string`

GetScheduledDeleteMsgDayList returns the ScheduledDeleteMsgDayList field if non-nil, zero value otherwise.

### GetScheduledDeleteMsgDayListOk

`func (o *MsgVpnDistributedCache) GetScheduledDeleteMsgDayListOk() (*string, bool)`

GetScheduledDeleteMsgDayListOk returns a tuple with the ScheduledDeleteMsgDayList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDeleteMsgDayList

`func (o *MsgVpnDistributedCache) SetScheduledDeleteMsgDayList(v string)`

SetScheduledDeleteMsgDayList sets ScheduledDeleteMsgDayList field to given value.

### HasScheduledDeleteMsgDayList

`func (o *MsgVpnDistributedCache) HasScheduledDeleteMsgDayList() bool`

HasScheduledDeleteMsgDayList returns a boolean if a field has been set.

### GetScheduledDeleteMsgTimeList

`func (o *MsgVpnDistributedCache) GetScheduledDeleteMsgTimeList() string`

GetScheduledDeleteMsgTimeList returns the ScheduledDeleteMsgTimeList field if non-nil, zero value otherwise.

### GetScheduledDeleteMsgTimeListOk

`func (o *MsgVpnDistributedCache) GetScheduledDeleteMsgTimeListOk() (*string, bool)`

GetScheduledDeleteMsgTimeListOk returns a tuple with the ScheduledDeleteMsgTimeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDeleteMsgTimeList

`func (o *MsgVpnDistributedCache) SetScheduledDeleteMsgTimeList(v string)`

SetScheduledDeleteMsgTimeList sets ScheduledDeleteMsgTimeList field to given value.

### HasScheduledDeleteMsgTimeList

`func (o *MsgVpnDistributedCache) HasScheduledDeleteMsgTimeList() bool`

HasScheduledDeleteMsgTimeList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


