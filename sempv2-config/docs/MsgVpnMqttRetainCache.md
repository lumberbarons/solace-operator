# MsgVpnMqttRetainCache

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CacheName** | Pointer to **string** | The name of the MQTT Retain Cache. | [optional] 
**Enabled** | Pointer to **bool** | Enable or disable this MQTT Retain Cache. When the cache is disabled, neither retain messages nor retain requests will be delivered by the cache. However, live retain messages will continue to be delivered to currently connected MQTT clients. The default value is &#x60;false&#x60;. | [optional] 
**MsgLifetime** | Pointer to **int64** | The message lifetime, in seconds. If a message remains cached for the duration of its lifetime, the cache will remove the message. A lifetime of 0 results in the message being retained indefinitely. The default value is &#x60;0&#x60;. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


