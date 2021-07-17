# MsgVpnDistributedCacheClusterInstanceRemoteTopic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheName** | Pointer to **string** | The name of the Distributed Cache. | [optional] 
**ClusterName** | Pointer to **string** | The name of the Cache Cluster. | [optional] 
**GlobalTopic** | Pointer to **bool** | Indicates whether the type of the remote Topic is global. | [optional] 
**HomeClusterName** | Pointer to **string** | The name of the remote Home Cache Cluster. | [optional] 
**InstanceName** | Pointer to **string** | The name of the Cache Instance. | [optional] 
**MsgCount** | Pointer to **int32** | The number of messages cached for the remote Topic. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**Suspect** | Pointer to **bool** | Indicates whether the remote Topic is suspect due to the remote Home Cache Cluster being in the lost message state. | [optional] 
**Topic** | Pointer to **string** | The value of the remote Topic. | [optional] 

## Methods

### NewMsgVpnDistributedCacheClusterInstanceRemoteTopic

`func NewMsgVpnDistributedCacheClusterInstanceRemoteTopic() *MsgVpnDistributedCacheClusterInstanceRemoteTopic`

NewMsgVpnDistributedCacheClusterInstanceRemoteTopic instantiates a new MsgVpnDistributedCacheClusterInstanceRemoteTopic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnDistributedCacheClusterInstanceRemoteTopicWithDefaults

`func NewMsgVpnDistributedCacheClusterInstanceRemoteTopicWithDefaults() *MsgVpnDistributedCacheClusterInstanceRemoteTopic`

NewMsgVpnDistributedCacheClusterInstanceRemoteTopicWithDefaults instantiates a new MsgVpnDistributedCacheClusterInstanceRemoteTopic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCacheName

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) GetCacheName() string`

GetCacheName returns the CacheName field if non-nil, zero value otherwise.

### GetCacheNameOk

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) GetCacheNameOk() (*string, bool)`

GetCacheNameOk returns a tuple with the CacheName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheName

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) SetCacheName(v string)`

SetCacheName sets CacheName field to given value.

### HasCacheName

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) HasCacheName() bool`

HasCacheName returns a boolean if a field has been set.

### GetClusterName

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetGlobalTopic

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) GetGlobalTopic() bool`

GetGlobalTopic returns the GlobalTopic field if non-nil, zero value otherwise.

### GetGlobalTopicOk

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) GetGlobalTopicOk() (*bool, bool)`

GetGlobalTopicOk returns a tuple with the GlobalTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalTopic

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) SetGlobalTopic(v bool)`

SetGlobalTopic sets GlobalTopic field to given value.

### HasGlobalTopic

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) HasGlobalTopic() bool`

HasGlobalTopic returns a boolean if a field has been set.

### GetHomeClusterName

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) GetHomeClusterName() string`

GetHomeClusterName returns the HomeClusterName field if non-nil, zero value otherwise.

### GetHomeClusterNameOk

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) GetHomeClusterNameOk() (*string, bool)`

GetHomeClusterNameOk returns a tuple with the HomeClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeClusterName

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) SetHomeClusterName(v string)`

SetHomeClusterName sets HomeClusterName field to given value.

### HasHomeClusterName

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) HasHomeClusterName() bool`

HasHomeClusterName returns a boolean if a field has been set.

### GetInstanceName

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### GetMsgCount

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) GetMsgCount() int32`

GetMsgCount returns the MsgCount field if non-nil, zero value otherwise.

### GetMsgCountOk

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) GetMsgCountOk() (*int32, bool)`

GetMsgCountOk returns a tuple with the MsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgCount

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) SetMsgCount(v int32)`

SetMsgCount sets MsgCount field to given value.

### HasMsgCount

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) HasMsgCount() bool`

HasMsgCount returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetSuspect

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) GetSuspect() bool`

GetSuspect returns the Suspect field if non-nil, zero value otherwise.

### GetSuspectOk

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) GetSuspectOk() (*bool, bool)`

GetSuspectOk returns a tuple with the Suspect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspect

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) SetSuspect(v bool)`

SetSuspect sets Suspect field to given value.

### HasSuspect

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) HasSuspect() bool`

HasSuspect returns a boolean if a field has been set.

### GetTopic

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *MsgVpnDistributedCacheClusterInstanceRemoteTopic) HasTopic() bool`

HasTopic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


