# MsgVpnReplayLogMsg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachmentSize** | Pointer to **int64** | The size of the message attachment, in bytes (B). | [optional] 
**ContentSize** | Pointer to **int64** | The size of the message content, in bytes (B). | [optional] 
**DmqEligible** | Pointer to **bool** | Indicates whether the message is eligible for the Dead Message Queue (DMQ). | [optional] 
**MsgId** | Pointer to **int64** | The identifier (ID) of the message. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**Priority** | Pointer to **int32** | The priority level of the message. | [optional] 
**PublisherId** | Pointer to **int64** | The identifier (ID) of the message publisher. | [optional] 
**ReplayLogName** | Pointer to **string** | The name of the Replay Log. | [optional] 
**ReplicationGroupMsgId** | Pointer to **string** | An ID that uniquely identifies this Message within this replication group. Available since 2.21. | [optional] 
**SequenceNumber** | Pointer to **int64** | The sequence number assigned to the message. Applicable only to messages received on sequenced topics. | [optional] 
**SpooledTime** | Pointer to **int32** | The timestamp of when the message was spooled in the Replay Log. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] 

## Methods

### NewMsgVpnReplayLogMsg

`func NewMsgVpnReplayLogMsg() *MsgVpnReplayLogMsg`

NewMsgVpnReplayLogMsg instantiates a new MsgVpnReplayLogMsg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnReplayLogMsgWithDefaults

`func NewMsgVpnReplayLogMsgWithDefaults() *MsgVpnReplayLogMsg`

NewMsgVpnReplayLogMsgWithDefaults instantiates a new MsgVpnReplayLogMsg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachmentSize

`func (o *MsgVpnReplayLogMsg) GetAttachmentSize() int64`

GetAttachmentSize returns the AttachmentSize field if non-nil, zero value otherwise.

### GetAttachmentSizeOk

`func (o *MsgVpnReplayLogMsg) GetAttachmentSizeOk() (*int64, bool)`

GetAttachmentSizeOk returns a tuple with the AttachmentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentSize

`func (o *MsgVpnReplayLogMsg) SetAttachmentSize(v int64)`

SetAttachmentSize sets AttachmentSize field to given value.

### HasAttachmentSize

`func (o *MsgVpnReplayLogMsg) HasAttachmentSize() bool`

HasAttachmentSize returns a boolean if a field has been set.

### GetContentSize

`func (o *MsgVpnReplayLogMsg) GetContentSize() int64`

GetContentSize returns the ContentSize field if non-nil, zero value otherwise.

### GetContentSizeOk

`func (o *MsgVpnReplayLogMsg) GetContentSizeOk() (*int64, bool)`

GetContentSizeOk returns a tuple with the ContentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSize

`func (o *MsgVpnReplayLogMsg) SetContentSize(v int64)`

SetContentSize sets ContentSize field to given value.

### HasContentSize

`func (o *MsgVpnReplayLogMsg) HasContentSize() bool`

HasContentSize returns a boolean if a field has been set.

### GetDmqEligible

`func (o *MsgVpnReplayLogMsg) GetDmqEligible() bool`

GetDmqEligible returns the DmqEligible field if non-nil, zero value otherwise.

### GetDmqEligibleOk

`func (o *MsgVpnReplayLogMsg) GetDmqEligibleOk() (*bool, bool)`

GetDmqEligibleOk returns a tuple with the DmqEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmqEligible

`func (o *MsgVpnReplayLogMsg) SetDmqEligible(v bool)`

SetDmqEligible sets DmqEligible field to given value.

### HasDmqEligible

`func (o *MsgVpnReplayLogMsg) HasDmqEligible() bool`

HasDmqEligible returns a boolean if a field has been set.

### GetMsgId

`func (o *MsgVpnReplayLogMsg) GetMsgId() int64`

GetMsgId returns the MsgId field if non-nil, zero value otherwise.

### GetMsgIdOk

`func (o *MsgVpnReplayLogMsg) GetMsgIdOk() (*int64, bool)`

GetMsgIdOk returns a tuple with the MsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgId

`func (o *MsgVpnReplayLogMsg) SetMsgId(v int64)`

SetMsgId sets MsgId field to given value.

### HasMsgId

`func (o *MsgVpnReplayLogMsg) HasMsgId() bool`

HasMsgId returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnReplayLogMsg) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnReplayLogMsg) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnReplayLogMsg) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnReplayLogMsg) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetPriority

`func (o *MsgVpnReplayLogMsg) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MsgVpnReplayLogMsg) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *MsgVpnReplayLogMsg) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *MsgVpnReplayLogMsg) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPublisherId

`func (o *MsgVpnReplayLogMsg) GetPublisherId() int64`

GetPublisherId returns the PublisherId field if non-nil, zero value otherwise.

### GetPublisherIdOk

`func (o *MsgVpnReplayLogMsg) GetPublisherIdOk() (*int64, bool)`

GetPublisherIdOk returns a tuple with the PublisherId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherId

`func (o *MsgVpnReplayLogMsg) SetPublisherId(v int64)`

SetPublisherId sets PublisherId field to given value.

### HasPublisherId

`func (o *MsgVpnReplayLogMsg) HasPublisherId() bool`

HasPublisherId returns a boolean if a field has been set.

### GetReplayLogName

`func (o *MsgVpnReplayLogMsg) GetReplayLogName() string`

GetReplayLogName returns the ReplayLogName field if non-nil, zero value otherwise.

### GetReplayLogNameOk

`func (o *MsgVpnReplayLogMsg) GetReplayLogNameOk() (*string, bool)`

GetReplayLogNameOk returns a tuple with the ReplayLogName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplayLogName

`func (o *MsgVpnReplayLogMsg) SetReplayLogName(v string)`

SetReplayLogName sets ReplayLogName field to given value.

### HasReplayLogName

`func (o *MsgVpnReplayLogMsg) HasReplayLogName() bool`

HasReplayLogName returns a boolean if a field has been set.

### GetReplicationGroupMsgId

`func (o *MsgVpnReplayLogMsg) GetReplicationGroupMsgId() string`

GetReplicationGroupMsgId returns the ReplicationGroupMsgId field if non-nil, zero value otherwise.

### GetReplicationGroupMsgIdOk

`func (o *MsgVpnReplayLogMsg) GetReplicationGroupMsgIdOk() (*string, bool)`

GetReplicationGroupMsgIdOk returns a tuple with the ReplicationGroupMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationGroupMsgId

`func (o *MsgVpnReplayLogMsg) SetReplicationGroupMsgId(v string)`

SetReplicationGroupMsgId sets ReplicationGroupMsgId field to given value.

### HasReplicationGroupMsgId

`func (o *MsgVpnReplayLogMsg) HasReplicationGroupMsgId() bool`

HasReplicationGroupMsgId returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *MsgVpnReplayLogMsg) GetSequenceNumber() int64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *MsgVpnReplayLogMsg) GetSequenceNumberOk() (*int64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *MsgVpnReplayLogMsg) SetSequenceNumber(v int64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *MsgVpnReplayLogMsg) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetSpooledTime

`func (o *MsgVpnReplayLogMsg) GetSpooledTime() int32`

GetSpooledTime returns the SpooledTime field if non-nil, zero value otherwise.

### GetSpooledTimeOk

`func (o *MsgVpnReplayLogMsg) GetSpooledTimeOk() (*int32, bool)`

GetSpooledTimeOk returns a tuple with the SpooledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpooledTime

`func (o *MsgVpnReplayLogMsg) SetSpooledTime(v int32)`

SetSpooledTime sets SpooledTime field to given value.

### HasSpooledTime

`func (o *MsgVpnReplayLogMsg) HasSpooledTime() bool`

HasSpooledTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


