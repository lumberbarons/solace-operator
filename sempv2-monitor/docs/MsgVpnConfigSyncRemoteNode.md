# MsgVpnConfigSyncRemoteNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastMsgRxTime** | Pointer to **int32** | The amount of time in seconds since the last message was received from the config sync table of the remote Message VPN. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**RemoteNodeName** | Pointer to **string** | The name of the Config Sync Remote Node. | [optional] 
**Role** | Pointer to **string** | The role of the config sync table of the remote Message VPN. The allowed values and their meaning are:  &lt;pre&gt; \&quot;unknown\&quot; - The role is unknown. \&quot;primary\&quot; - Acts as the primary source of config data. \&quot;replica\&quot; - Acts as a replica of the primary config data. &lt;/pre&gt;  | [optional] 
**Stale** | Pointer to **bool** | Indicates whether the config sync table of the remote Message VPN is stale. | [optional] 
**State** | Pointer to **string** | The state of the config sync table of the remote Message VPN. The allowed values and their meaning are:  &lt;pre&gt; \&quot;unknown\&quot; - The state is unknown. \&quot;in-sync\&quot; - The config data is synchronized between Message VPNs. \&quot;reconciling\&quot; - The config data is reconciling between Message VPNs. \&quot;blocked\&quot; - The config data is blocked from reconciling due to an error. \&quot;out-of-sync\&quot; - The config data is out of sync between Message VPNs. \&quot;down\&quot; - The state is down due to configuration. &lt;/pre&gt;  | [optional] 
**TimeInState** | Pointer to **int32** | The amount of time in seconds the config sync table of the remote Message VPN has been in the current state. | [optional] 

## Methods

### NewMsgVpnConfigSyncRemoteNode

`func NewMsgVpnConfigSyncRemoteNode() *MsgVpnConfigSyncRemoteNode`

NewMsgVpnConfigSyncRemoteNode instantiates a new MsgVpnConfigSyncRemoteNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnConfigSyncRemoteNodeWithDefaults

`func NewMsgVpnConfigSyncRemoteNodeWithDefaults() *MsgVpnConfigSyncRemoteNode`

NewMsgVpnConfigSyncRemoteNodeWithDefaults instantiates a new MsgVpnConfigSyncRemoteNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastMsgRxTime

`func (o *MsgVpnConfigSyncRemoteNode) GetLastMsgRxTime() int32`

GetLastMsgRxTime returns the LastMsgRxTime field if non-nil, zero value otherwise.

### GetLastMsgRxTimeOk

`func (o *MsgVpnConfigSyncRemoteNode) GetLastMsgRxTimeOk() (*int32, bool)`

GetLastMsgRxTimeOk returns a tuple with the LastMsgRxTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMsgRxTime

`func (o *MsgVpnConfigSyncRemoteNode) SetLastMsgRxTime(v int32)`

SetLastMsgRxTime sets LastMsgRxTime field to given value.

### HasLastMsgRxTime

`func (o *MsgVpnConfigSyncRemoteNode) HasLastMsgRxTime() bool`

HasLastMsgRxTime returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnConfigSyncRemoteNode) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnConfigSyncRemoteNode) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnConfigSyncRemoteNode) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnConfigSyncRemoteNode) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetRemoteNodeName

`func (o *MsgVpnConfigSyncRemoteNode) GetRemoteNodeName() string`

GetRemoteNodeName returns the RemoteNodeName field if non-nil, zero value otherwise.

### GetRemoteNodeNameOk

`func (o *MsgVpnConfigSyncRemoteNode) GetRemoteNodeNameOk() (*string, bool)`

GetRemoteNodeNameOk returns a tuple with the RemoteNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteNodeName

`func (o *MsgVpnConfigSyncRemoteNode) SetRemoteNodeName(v string)`

SetRemoteNodeName sets RemoteNodeName field to given value.

### HasRemoteNodeName

`func (o *MsgVpnConfigSyncRemoteNode) HasRemoteNodeName() bool`

HasRemoteNodeName returns a boolean if a field has been set.

### GetRole

`func (o *MsgVpnConfigSyncRemoteNode) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MsgVpnConfigSyncRemoteNode) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MsgVpnConfigSyncRemoteNode) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *MsgVpnConfigSyncRemoteNode) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStale

`func (o *MsgVpnConfigSyncRemoteNode) GetStale() bool`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *MsgVpnConfigSyncRemoteNode) GetStaleOk() (*bool, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *MsgVpnConfigSyncRemoteNode) SetStale(v bool)`

SetStale sets Stale field to given value.

### HasStale

`func (o *MsgVpnConfigSyncRemoteNode) HasStale() bool`

HasStale returns a boolean if a field has been set.

### GetState

`func (o *MsgVpnConfigSyncRemoteNode) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MsgVpnConfigSyncRemoteNode) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MsgVpnConfigSyncRemoteNode) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *MsgVpnConfigSyncRemoteNode) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTimeInState

`func (o *MsgVpnConfigSyncRemoteNode) GetTimeInState() int32`

GetTimeInState returns the TimeInState field if non-nil, zero value otherwise.

### GetTimeInStateOk

`func (o *MsgVpnConfigSyncRemoteNode) GetTimeInStateOk() (*int32, bool)`

GetTimeInStateOk returns a tuple with the TimeInState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeInState

`func (o *MsgVpnConfigSyncRemoteNode) SetTimeInState(v int32)`

SetTimeInState sets TimeInState field to given value.

### HasTimeInState

`func (o *MsgVpnConfigSyncRemoteNode) HasTimeInState() bool`

HasTimeInState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


