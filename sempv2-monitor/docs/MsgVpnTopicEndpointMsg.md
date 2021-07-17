# MsgVpnTopicEndpointMsg

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
**RedeliveryCount** | Pointer to **int32** | The number of times the Message has been redelivered. | [optional] 
**ReplicatedMateMsgId** | Pointer to **int64** | The Message identifier (ID) on the replication mate. Applicable only to replicated messages. | [optional] 
**ReplicationGroupMsgId** | Pointer to **string** | An ID that uniquely identifies this Message within this replication group. Available since 2.21. | [optional] 
**ReplicationState** | Pointer to **string** | The replication state of the Message. The allowed values and their meaning are:  &lt;pre&gt; \&quot;replicated\&quot; - The Message is replicated to the remote Message VPN. \&quot;not-replicated\&quot; - The Message is not being replicated to the remote Message VPN. \&quot;pending-replication\&quot; - The Message is queued for replication to the remote Message VPN. &lt;/pre&gt;  | [optional] 
**SpooledTime** | Pointer to **int32** | The timestamp of when the Message was spooled in the Topic Endpoint. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] 
**TopicEndpointName** | Pointer to **string** | The name of the Topic Endpoint. | [optional] 
**Undelivered** | Pointer to **bool** | Indicates whether delivery of the Message has never been attempted. | [optional] 

## Methods

### NewMsgVpnTopicEndpointMsg

`func NewMsgVpnTopicEndpointMsg() *MsgVpnTopicEndpointMsg`

NewMsgVpnTopicEndpointMsg instantiates a new MsgVpnTopicEndpointMsg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnTopicEndpointMsgWithDefaults

`func NewMsgVpnTopicEndpointMsgWithDefaults() *MsgVpnTopicEndpointMsg`

NewMsgVpnTopicEndpointMsgWithDefaults instantiates a new MsgVpnTopicEndpointMsg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachmentSize

`func (o *MsgVpnTopicEndpointMsg) GetAttachmentSize() int64`

GetAttachmentSize returns the AttachmentSize field if non-nil, zero value otherwise.

### GetAttachmentSizeOk

`func (o *MsgVpnTopicEndpointMsg) GetAttachmentSizeOk() (*int64, bool)`

GetAttachmentSizeOk returns a tuple with the AttachmentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentSize

`func (o *MsgVpnTopicEndpointMsg) SetAttachmentSize(v int64)`

SetAttachmentSize sets AttachmentSize field to given value.

### HasAttachmentSize

`func (o *MsgVpnTopicEndpointMsg) HasAttachmentSize() bool`

HasAttachmentSize returns a boolean if a field has been set.

### GetContentSize

`func (o *MsgVpnTopicEndpointMsg) GetContentSize() int64`

GetContentSize returns the ContentSize field if non-nil, zero value otherwise.

### GetContentSizeOk

`func (o *MsgVpnTopicEndpointMsg) GetContentSizeOk() (*int64, bool)`

GetContentSizeOk returns a tuple with the ContentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSize

`func (o *MsgVpnTopicEndpointMsg) SetContentSize(v int64)`

SetContentSize sets ContentSize field to given value.

### HasContentSize

`func (o *MsgVpnTopicEndpointMsg) HasContentSize() bool`

HasContentSize returns a boolean if a field has been set.

### GetDmqEligible

`func (o *MsgVpnTopicEndpointMsg) GetDmqEligible() bool`

GetDmqEligible returns the DmqEligible field if non-nil, zero value otherwise.

### GetDmqEligibleOk

`func (o *MsgVpnTopicEndpointMsg) GetDmqEligibleOk() (*bool, bool)`

GetDmqEligibleOk returns a tuple with the DmqEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmqEligible

`func (o *MsgVpnTopicEndpointMsg) SetDmqEligible(v bool)`

SetDmqEligible sets DmqEligible field to given value.

### HasDmqEligible

`func (o *MsgVpnTopicEndpointMsg) HasDmqEligible() bool`

HasDmqEligible returns a boolean if a field has been set.

### GetExpiryTime

`func (o *MsgVpnTopicEndpointMsg) GetExpiryTime() int32`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *MsgVpnTopicEndpointMsg) GetExpiryTimeOk() (*int32, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *MsgVpnTopicEndpointMsg) SetExpiryTime(v int32)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *MsgVpnTopicEndpointMsg) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.

### GetMsgId

`func (o *MsgVpnTopicEndpointMsg) GetMsgId() int64`

GetMsgId returns the MsgId field if non-nil, zero value otherwise.

### GetMsgIdOk

`func (o *MsgVpnTopicEndpointMsg) GetMsgIdOk() (*int64, bool)`

GetMsgIdOk returns a tuple with the MsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgId

`func (o *MsgVpnTopicEndpointMsg) SetMsgId(v int64)`

SetMsgId sets MsgId field to given value.

### HasMsgId

`func (o *MsgVpnTopicEndpointMsg) HasMsgId() bool`

HasMsgId returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnTopicEndpointMsg) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnTopicEndpointMsg) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnTopicEndpointMsg) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnTopicEndpointMsg) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetPriority

`func (o *MsgVpnTopicEndpointMsg) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MsgVpnTopicEndpointMsg) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *MsgVpnTopicEndpointMsg) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *MsgVpnTopicEndpointMsg) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPublisherId

`func (o *MsgVpnTopicEndpointMsg) GetPublisherId() int64`

GetPublisherId returns the PublisherId field if non-nil, zero value otherwise.

### GetPublisherIdOk

`func (o *MsgVpnTopicEndpointMsg) GetPublisherIdOk() (*int64, bool)`

GetPublisherIdOk returns a tuple with the PublisherId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherId

`func (o *MsgVpnTopicEndpointMsg) SetPublisherId(v int64)`

SetPublisherId sets PublisherId field to given value.

### HasPublisherId

`func (o *MsgVpnTopicEndpointMsg) HasPublisherId() bool`

HasPublisherId returns a boolean if a field has been set.

### GetRedeliveryCount

`func (o *MsgVpnTopicEndpointMsg) GetRedeliveryCount() int32`

GetRedeliveryCount returns the RedeliveryCount field if non-nil, zero value otherwise.

### GetRedeliveryCountOk

`func (o *MsgVpnTopicEndpointMsg) GetRedeliveryCountOk() (*int32, bool)`

GetRedeliveryCountOk returns a tuple with the RedeliveryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeliveryCount

`func (o *MsgVpnTopicEndpointMsg) SetRedeliveryCount(v int32)`

SetRedeliveryCount sets RedeliveryCount field to given value.

### HasRedeliveryCount

`func (o *MsgVpnTopicEndpointMsg) HasRedeliveryCount() bool`

HasRedeliveryCount returns a boolean if a field has been set.

### GetReplicatedMateMsgId

`func (o *MsgVpnTopicEndpointMsg) GetReplicatedMateMsgId() int64`

GetReplicatedMateMsgId returns the ReplicatedMateMsgId field if non-nil, zero value otherwise.

### GetReplicatedMateMsgIdOk

`func (o *MsgVpnTopicEndpointMsg) GetReplicatedMateMsgIdOk() (*int64, bool)`

GetReplicatedMateMsgIdOk returns a tuple with the ReplicatedMateMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicatedMateMsgId

`func (o *MsgVpnTopicEndpointMsg) SetReplicatedMateMsgId(v int64)`

SetReplicatedMateMsgId sets ReplicatedMateMsgId field to given value.

### HasReplicatedMateMsgId

`func (o *MsgVpnTopicEndpointMsg) HasReplicatedMateMsgId() bool`

HasReplicatedMateMsgId returns a boolean if a field has been set.

### GetReplicationGroupMsgId

`func (o *MsgVpnTopicEndpointMsg) GetReplicationGroupMsgId() string`

GetReplicationGroupMsgId returns the ReplicationGroupMsgId field if non-nil, zero value otherwise.

### GetReplicationGroupMsgIdOk

`func (o *MsgVpnTopicEndpointMsg) GetReplicationGroupMsgIdOk() (*string, bool)`

GetReplicationGroupMsgIdOk returns a tuple with the ReplicationGroupMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationGroupMsgId

`func (o *MsgVpnTopicEndpointMsg) SetReplicationGroupMsgId(v string)`

SetReplicationGroupMsgId sets ReplicationGroupMsgId field to given value.

### HasReplicationGroupMsgId

`func (o *MsgVpnTopicEndpointMsg) HasReplicationGroupMsgId() bool`

HasReplicationGroupMsgId returns a boolean if a field has been set.

### GetReplicationState

`func (o *MsgVpnTopicEndpointMsg) GetReplicationState() string`

GetReplicationState returns the ReplicationState field if non-nil, zero value otherwise.

### GetReplicationStateOk

`func (o *MsgVpnTopicEndpointMsg) GetReplicationStateOk() (*string, bool)`

GetReplicationStateOk returns a tuple with the ReplicationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationState

`func (o *MsgVpnTopicEndpointMsg) SetReplicationState(v string)`

SetReplicationState sets ReplicationState field to given value.

### HasReplicationState

`func (o *MsgVpnTopicEndpointMsg) HasReplicationState() bool`

HasReplicationState returns a boolean if a field has been set.

### GetSpooledTime

`func (o *MsgVpnTopicEndpointMsg) GetSpooledTime() int32`

GetSpooledTime returns the SpooledTime field if non-nil, zero value otherwise.

### GetSpooledTimeOk

`func (o *MsgVpnTopicEndpointMsg) GetSpooledTimeOk() (*int32, bool)`

GetSpooledTimeOk returns a tuple with the SpooledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpooledTime

`func (o *MsgVpnTopicEndpointMsg) SetSpooledTime(v int32)`

SetSpooledTime sets SpooledTime field to given value.

### HasSpooledTime

`func (o *MsgVpnTopicEndpointMsg) HasSpooledTime() bool`

HasSpooledTime returns a boolean if a field has been set.

### GetTopicEndpointName

`func (o *MsgVpnTopicEndpointMsg) GetTopicEndpointName() string`

GetTopicEndpointName returns the TopicEndpointName field if non-nil, zero value otherwise.

### GetTopicEndpointNameOk

`func (o *MsgVpnTopicEndpointMsg) GetTopicEndpointNameOk() (*string, bool)`

GetTopicEndpointNameOk returns a tuple with the TopicEndpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicEndpointName

`func (o *MsgVpnTopicEndpointMsg) SetTopicEndpointName(v string)`

SetTopicEndpointName sets TopicEndpointName field to given value.

### HasTopicEndpointName

`func (o *MsgVpnTopicEndpointMsg) HasTopicEndpointName() bool`

HasTopicEndpointName returns a boolean if a field has been set.

### GetUndelivered

`func (o *MsgVpnTopicEndpointMsg) GetUndelivered() bool`

GetUndelivered returns the Undelivered field if non-nil, zero value otherwise.

### GetUndeliveredOk

`func (o *MsgVpnTopicEndpointMsg) GetUndeliveredOk() (*bool, bool)`

GetUndeliveredOk returns a tuple with the Undelivered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUndelivered

`func (o *MsgVpnTopicEndpointMsg) SetUndelivered(v bool)`

SetUndelivered sets Undelivered field to given value.

### HasUndelivered

`func (o *MsgVpnTopicEndpointMsg) HasUndelivered() bool`

HasUndelivered returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


