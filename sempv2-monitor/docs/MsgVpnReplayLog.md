# MsgVpnReplayLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EgressEnabled** | Pointer to **bool** | Indicates whether the transmission of messages from the Replay Log is enabled. | [optional] 
**IngressEnabled** | Pointer to **bool** | Indicates whether the reception of messages to the Replay Log is enabled. | [optional] 
**MaxSpoolUsage** | Pointer to **int64** | The maximum spool usage allowed by the Replay Log, in megabytes (MB). If this limit is exceeded, old messages will be trimmed. | [optional] 
**MsgSpoolUsage** | Pointer to **int64** | The spool usage of the Replay Log, in bytes (B). | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**ReplayLogName** | Pointer to **string** | The name of the Replay Log. | [optional] 

## Methods

### NewMsgVpnReplayLog

`func NewMsgVpnReplayLog() *MsgVpnReplayLog`

NewMsgVpnReplayLog instantiates a new MsgVpnReplayLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnReplayLogWithDefaults

`func NewMsgVpnReplayLogWithDefaults() *MsgVpnReplayLog`

NewMsgVpnReplayLogWithDefaults instantiates a new MsgVpnReplayLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEgressEnabled

`func (o *MsgVpnReplayLog) GetEgressEnabled() bool`

GetEgressEnabled returns the EgressEnabled field if non-nil, zero value otherwise.

### GetEgressEnabledOk

`func (o *MsgVpnReplayLog) GetEgressEnabledOk() (*bool, bool)`

GetEgressEnabledOk returns a tuple with the EgressEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressEnabled

`func (o *MsgVpnReplayLog) SetEgressEnabled(v bool)`

SetEgressEnabled sets EgressEnabled field to given value.

### HasEgressEnabled

`func (o *MsgVpnReplayLog) HasEgressEnabled() bool`

HasEgressEnabled returns a boolean if a field has been set.

### GetIngressEnabled

`func (o *MsgVpnReplayLog) GetIngressEnabled() bool`

GetIngressEnabled returns the IngressEnabled field if non-nil, zero value otherwise.

### GetIngressEnabledOk

`func (o *MsgVpnReplayLog) GetIngressEnabledOk() (*bool, bool)`

GetIngressEnabledOk returns a tuple with the IngressEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressEnabled

`func (o *MsgVpnReplayLog) SetIngressEnabled(v bool)`

SetIngressEnabled sets IngressEnabled field to given value.

### HasIngressEnabled

`func (o *MsgVpnReplayLog) HasIngressEnabled() bool`

HasIngressEnabled returns a boolean if a field has been set.

### GetMaxSpoolUsage

`func (o *MsgVpnReplayLog) GetMaxSpoolUsage() int64`

GetMaxSpoolUsage returns the MaxSpoolUsage field if non-nil, zero value otherwise.

### GetMaxSpoolUsageOk

`func (o *MsgVpnReplayLog) GetMaxSpoolUsageOk() (*int64, bool)`

GetMaxSpoolUsageOk returns a tuple with the MaxSpoolUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpoolUsage

`func (o *MsgVpnReplayLog) SetMaxSpoolUsage(v int64)`

SetMaxSpoolUsage sets MaxSpoolUsage field to given value.

### HasMaxSpoolUsage

`func (o *MsgVpnReplayLog) HasMaxSpoolUsage() bool`

HasMaxSpoolUsage returns a boolean if a field has been set.

### GetMsgSpoolUsage

`func (o *MsgVpnReplayLog) GetMsgSpoolUsage() int64`

GetMsgSpoolUsage returns the MsgSpoolUsage field if non-nil, zero value otherwise.

### GetMsgSpoolUsageOk

`func (o *MsgVpnReplayLog) GetMsgSpoolUsageOk() (*int64, bool)`

GetMsgSpoolUsageOk returns a tuple with the MsgSpoolUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgSpoolUsage

`func (o *MsgVpnReplayLog) SetMsgSpoolUsage(v int64)`

SetMsgSpoolUsage sets MsgSpoolUsage field to given value.

### HasMsgSpoolUsage

`func (o *MsgVpnReplayLog) HasMsgSpoolUsage() bool`

HasMsgSpoolUsage returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnReplayLog) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnReplayLog) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnReplayLog) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnReplayLog) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetReplayLogName

`func (o *MsgVpnReplayLog) GetReplayLogName() string`

GetReplayLogName returns the ReplayLogName field if non-nil, zero value otherwise.

### GetReplayLogNameOk

`func (o *MsgVpnReplayLog) GetReplayLogNameOk() (*string, bool)`

GetReplayLogNameOk returns a tuple with the ReplayLogName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayLogName

`func (o *MsgVpnReplayLog) SetReplayLogName(v string)`

SetReplayLogName sets ReplayLogName field to given value.

### HasReplayLogName

`func (o *MsgVpnReplayLog) HasReplayLogName() bool`

HasReplayLogName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


