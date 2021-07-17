# MsgVpnQueueMsg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachmentSize** | Pointer to **int64** | The size of the Message attachment, in bytes (B). | [optional] 
**ContentSize** | Pointer to **int64** | The size of the Message content, in bytes (B). | [optional] 
**DmqEligible** | Pointer to **bool** | Indicates whether the Message is eligible for the Dead Message Queue (DMQ). | [optional] 
**ExpiryTime** | Pointer to **int32** | The timestamp of when the Message expires. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] 
**MsgId** | Pointer to **int64** | The identifier (ID) of the Message. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**Priority** | Pointer to **int32** | The priority level of the Message, from 9 (highest) to 0 (lowest). | [optional] 
**PublisherId** | Pointer to **int64** | The identifier (ID) of the Message publisher. | [optional] 
**QueueName** | Pointer to **string** | The name of the Queue. | [optional] 
**RedeliveryCount** | Pointer to **int32** | The number of times the Message has been redelivered. | [optional] 
**ReplicatedMateMsgId** | Pointer to **int64** | The Message identifier (ID) on the replication mate. Applicable only to replicated messages. | [optional] 
**ReplicationGroupMsgId** | Pointer to **string** | An ID that uniquely identifies this Message within this replication group. Available since 2.21. | [optional] 
**ReplicationState** | Pointer to **string** | The replication state of the Message. The allowed values and their meaning are:  &lt;pre&gt; \&quot;replicated\&quot; - The Message is replicated to the remote Message VPN. \&quot;not-replicated\&quot; - The Message is not being replicated to the remote Message VPN. \&quot;pending-replication\&quot; - The Message is queued for replication to the remote Message VPN. &lt;/pre&gt;  | [optional] 
**SpooledTime** | Pointer to **int32** | The timestamp of when the Message was spooled in the Queue. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] 
**Undelivered** | Pointer to **bool** | Indicates whether delivery of the Message has never been attempted. | [optional] 

## Methods

### NewMsgVpnQueueMsg

`func NewMsgVpnQueueMsg() *MsgVpnQueueMsg`

NewMsgVpnQueueMsg instantiates a new MsgVpnQueueMsg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnQueueMsgWithDefaults

`func NewMsgVpnQueueMsgWithDefaults() *MsgVpnQueueMsg`

NewMsgVpnQueueMsgWithDefaults instantiates a new MsgVpnQueueMsg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachmentSize

`func (o *MsgVpnQueueMsg) GetAttachmentSize() int64`

GetAttachmentSize returns the AttachmentSize field if non-nil, zero value otherwise.

### GetAttachmentSizeOk

`func (o *MsgVpnQueueMsg) GetAttachmentSizeOk() (*int64, bool)`

GetAttachmentSizeOk returns a tuple with the AttachmentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentSize

`func (o *MsgVpnQueueMsg) SetAttachmentSize(v int64)`

SetAttachmentSize sets AttachmentSize field to given value.

### HasAttachmentSize

`func (o *MsgVpnQueueMsg) HasAttachmentSize() bool`

HasAttachmentSize returns a boolean if a field has been set.

### GetContentSize

`func (o *MsgVpnQueueMsg) GetContentSize() int64`

GetContentSize returns the ContentSize field if non-nil, zero value otherwise.

### GetContentSizeOk

`func (o *MsgVpnQueueMsg) GetContentSizeOk() (*int64, bool)`

GetContentSizeOk returns a tuple with the ContentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSize

`func (o *MsgVpnQueueMsg) SetContentSize(v int64)`

SetContentSize sets ContentSize field to given value.

### HasContentSize

`func (o *MsgVpnQueueMsg) HasContentSize() bool`

HasContentSize returns a boolean if a field has been set.

### GetDmqEligible

`func (o *MsgVpnQueueMsg) GetDmqEligible() bool`

GetDmqEligible returns the DmqEligible field if non-nil, zero value otherwise.

### GetDmqEligibleOk

`func (o *MsgVpnQueueMsg) GetDmqEligibleOk() (*bool, bool)`

GetDmqEligibleOk returns a tuple with the DmqEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmqEligible

`func (o *MsgVpnQueueMsg) SetDmqEligible(v bool)`

SetDmqEligible sets DmqEligible field to given value.

### HasDmqEligible

`func (o *MsgVpnQueueMsg) HasDmqEligible() bool`

HasDmqEligible returns a boolean if a field has been set.

### GetExpiryTime

`func (o *MsgVpnQueueMsg) GetExpiryTime() int32`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *MsgVpnQueueMsg) GetExpiryTimeOk() (*int32, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *MsgVpnQueueMsg) SetExpiryTime(v int32)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *MsgVpnQueueMsg) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.

### GetMsgId

`func (o *MsgVpnQueueMsg) GetMsgId() int64`

GetMsgId returns the MsgId field if non-nil, zero value otherwise.

### GetMsgIdOk

`func (o *MsgVpnQueueMsg) GetMsgIdOk() (*int64, bool)`

GetMsgIdOk returns a tuple with the MsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgId

`func (o *MsgVpnQueueMsg) SetMsgId(v int64)`

SetMsgId sets MsgId field to given value.

### HasMsgId

`func (o *MsgVpnQueueMsg) HasMsgId() bool`

HasMsgId returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnQueueMsg) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnQueueMsg) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnQueueMsg) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnQueueMsg) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetPriority

`func (o *MsgVpnQueueMsg) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MsgVpnQueueMsg) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *MsgVpnQueueMsg) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *MsgVpnQueueMsg) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPublisherId

`func (o *MsgVpnQueueMsg) GetPublisherId() int64`

GetPublisherId returns the PublisherId field if non-nil, zero value otherwise.

### GetPublisherIdOk

`func (o *MsgVpnQueueMsg) GetPublisherIdOk() (*int64, bool)`

GetPublisherIdOk returns a tuple with the PublisherId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherId

`func (o *MsgVpnQueueMsg) SetPublisherId(v int64)`

SetPublisherId sets PublisherId field to given value.

### HasPublisherId

`func (o *MsgVpnQueueMsg) HasPublisherId() bool`

HasPublisherId returns a boolean if a field has been set.

### GetQueueName

`func (o *MsgVpnQueueMsg) GetQueueName() string`

GetQueueName returns the QueueName field if non-nil, zero value otherwise.

### GetQueueNameOk

`func (o *MsgVpnQueueMsg) GetQueueNameOk() (*string, bool)`

GetQueueNameOk returns a tuple with the QueueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueName

`func (o *MsgVpnQueueMsg) SetQueueName(v string)`

SetQueueName sets QueueName field to given value.

### HasQueueName

`func (o *MsgVpnQueueMsg) HasQueueName() bool`

HasQueueName returns a boolean if a field has been set.

### GetRedeliveryCount

`func (o *MsgVpnQueueMsg) GetRedeliveryCount() int32`

GetRedeliveryCount returns the RedeliveryCount field if non-nil, zero value otherwise.

### GetRedeliveryCountOk

`func (o *MsgVpnQueueMsg) GetRedeliveryCountOk() (*int32, bool)`

GetRedeliveryCountOk returns a tuple with the RedeliveryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeliveryCount

`func (o *MsgVpnQueueMsg) SetRedeliveryCount(v int32)`

SetRedeliveryCount sets RedeliveryCount field to given value.

### HasRedeliveryCount

`func (o *MsgVpnQueueMsg) HasRedeliveryCount() bool`

HasRedeliveryCount returns a boolean if a field has been set.

### GetReplicatedMateMsgId

`func (o *MsgVpnQueueMsg) GetReplicatedMateMsgId() int64`

GetReplicatedMateMsgId returns the ReplicatedMateMsgId field if non-nil, zero value otherwise.

### GetReplicatedMateMsgIdOk

`func (o *MsgVpnQueueMsg) GetReplicatedMateMsgIdOk() (*int64, bool)`

GetReplicatedMateMsgIdOk returns a tuple with the ReplicatedMateMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicatedMateMsgId

`func (o *MsgVpnQueueMsg) SetReplicatedMateMsgId(v int64)`

SetReplicatedMateMsgId sets ReplicatedMateMsgId field to given value.

### HasReplicatedMateMsgId

`func (o *MsgVpnQueueMsg) HasReplicatedMateMsgId() bool`

HasReplicatedMateMsgId returns a boolean if a field has been set.

### GetReplicationGroupMsgId

`func (o *MsgVpnQueueMsg) GetReplicationGroupMsgId() string`

GetReplicationGroupMsgId returns the ReplicationGroupMsgId field if non-nil, zero value otherwise.

### GetReplicationGroupMsgIdOk

`func (o *MsgVpnQueueMsg) GetReplicationGroupMsgIdOk() (*string, bool)`

GetReplicationGroupMsgIdOk returns a tuple with the ReplicationGroupMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationGroupMsgId

`func (o *MsgVpnQueueMsg) SetReplicationGroupMsgId(v string)`

SetReplicationGroupMsgId sets ReplicationGroupMsgId field to given value.

### HasReplicationGroupMsgId

`func (o *MsgVpnQueueMsg) HasReplicationGroupMsgId() bool`

HasReplicationGroupMsgId returns a boolean if a field has been set.

### GetReplicationState

`func (o *MsgVpnQueueMsg) GetReplicationState() string`

GetReplicationState returns the ReplicationState field if non-nil, zero value otherwise.

### GetReplicationStateOk

`func (o *MsgVpnQueueMsg) GetReplicationStateOk() (*string, bool)`

GetReplicationStateOk returns a tuple with the ReplicationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationState

`func (o *MsgVpnQueueMsg) SetReplicationState(v string)`

SetReplicationState sets ReplicationState field to given value.

### HasReplicationState

`func (o *MsgVpnQueueMsg) HasReplicationState() bool`

HasReplicationState returns a boolean if a field has been set.

### GetSpooledTime

`func (o *MsgVpnQueueMsg) GetSpooledTime() int32`

GetSpooledTime returns the SpooledTime field if non-nil, zero value otherwise.

### GetSpooledTimeOk

`func (o *MsgVpnQueueMsg) GetSpooledTimeOk() (*int32, bool)`

GetSpooledTimeOk returns a tuple with the SpooledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpooledTime

`func (o *MsgVpnQueueMsg) SetSpooledTime(v int32)`

SetSpooledTime sets SpooledTime field to given value.

### HasSpooledTime

`func (o *MsgVpnQueueMsg) HasSpooledTime() bool`

HasSpooledTime returns a boolean if a field has been set.

### GetUndelivered

`func (o *MsgVpnQueueMsg) GetUndelivered() bool`

GetUndelivered returns the Undelivered field if non-nil, zero value otherwise.

### GetUndeliveredOk

`func (o *MsgVpnQueueMsg) GetUndeliveredOk() (*bool, bool)`

GetUndeliveredOk returns a tuple with the Undelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUndelivered

`func (o *MsgVpnQueueMsg) SetUndelivered(v bool)`

SetUndelivered sets Undelivered field to given value.

### HasUndelivered

`func (o *MsgVpnQueueMsg) HasUndelivered() bool`

HasUndelivered returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


