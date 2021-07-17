# MsgVpnDistributedCacheClusterInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoStartEnabled** | Pointer to **bool** | Enable or disable auto-start for the Cache Instance. When enabled, the Cache Instance will automatically attempt to transition from the Stopped operational state to Up whenever it restarts or reconnects to the message broker. The default value is &#x60;false&#x60;. | [optional] 
**CacheName** | Pointer to **string** | The name of the Distributed Cache. | [optional] 
**ClusterName** | Pointer to **string** | The name of the Cache Cluster. | [optional] 
**Enabled** | Pointer to **bool** | Enable or disable the Cache Instance. The default value is &#x60;false&#x60;. | [optional] 
**InstanceName** | Pointer to **string** | The name of the Cache Instance. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**StopOnLostMsgEnabled** | Pointer to **bool** | Enable or disable stop-on-lost-message for the Cache Instance. When enabled, the Cache Instance will transition to the stopped operational state upon losing a message. When stopped, it cannot accept or respond to cache requests, but continues to cache messages. The default value is &#x60;true&#x60;. | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


