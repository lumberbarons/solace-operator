# MsgVpnReplicatedTopic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**ReplicatedTopic** | Pointer to **string** | The topic for applying replication. Published messages matching this topic will be replicated to the standby site. | [optional] 
**ReplicationMode** | Pointer to **string** | The replication mode for the Replicated Topic. The allowed values and their meaning are:  &lt;pre&gt; \&quot;sync\&quot; - Messages are acknowledged when replicated (spooled remotely). \&quot;async\&quot; - Messages are acknowledged when pending replication (spooled locally). &lt;/pre&gt;  | [optional] 

## Methods

### NewMsgVpnReplicatedTopic

`func NewMsgVpnReplicatedTopic() *MsgVpnReplicatedTopic`

NewMsgVpnReplicatedTopic instantiates a new MsgVpnReplicatedTopic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnReplicatedTopicWithDefaults

`func NewMsgVpnReplicatedTopicWithDefaults() *MsgVpnReplicatedTopic`

NewMsgVpnReplicatedTopicWithDefaults instantiates a new MsgVpnReplicatedTopic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsgVpnName

`func (o *MsgVpnReplicatedTopic) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnReplicatedTopic) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnReplicatedTopic) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnReplicatedTopic) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetReplicatedTopic

`func (o *MsgVpnReplicatedTopic) GetReplicatedTopic() string`

GetReplicatedTopic returns the ReplicatedTopic field if non-nil, zero value otherwise.

### GetReplicatedTopicOk

`func (o *MsgVpnReplicatedTopic) GetReplicatedTopicOk() (*string, bool)`

GetReplicatedTopicOk returns a tuple with the ReplicatedTopic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicatedTopic

`func (o *MsgVpnReplicatedTopic) SetReplicatedTopic(v string)`

SetReplicatedTopic sets ReplicatedTopic field to given value.

### HasReplicatedTopic

`func (o *MsgVpnReplicatedTopic) HasReplicatedTopic() bool`

HasReplicatedTopic returns a boolean if a field has been set.

### GetReplicationMode

`func (o *MsgVpnReplicatedTopic) GetReplicationMode() string`

GetReplicationMode returns the ReplicationMode field if non-nil, zero value otherwise.

### GetReplicationModeOk

`func (o *MsgVpnReplicatedTopic) GetReplicationModeOk() (*string, bool)`

GetReplicationModeOk returns a tuple with the ReplicationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationMode

`func (o *MsgVpnReplicatedTopic) SetReplicationMode(v string)`

SetReplicationMode sets ReplicationMode field to given value.

### HasReplicationMode

`func (o *MsgVpnReplicatedTopic) HasReplicationMode() bool`

HasReplicationMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


