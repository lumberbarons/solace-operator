# MsgVpnQueuePriority

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**Priority** | Pointer to **int32** | The level of the Priority, from 9 (highest) to 0 (lowest). | [optional] 
**QueueName** | Pointer to **string** | The name of the Queue. | [optional] 
**SpooledByteCount** | Pointer to **int64** | The amount of guaranteed messages at this Priority spooled by the Queue, in bytes (B). | [optional] 
**SpooledMsgCount** | Pointer to **int64** | The number of guaranteed messages at this Priority spooled by the Queue. | [optional] 

## Methods

### NewMsgVpnQueuePriority

`func NewMsgVpnQueuePriority() *MsgVpnQueuePriority`

NewMsgVpnQueuePriority instantiates a new MsgVpnQueuePriority object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnQueuePriorityWithDefaults

`func NewMsgVpnQueuePriorityWithDefaults() *MsgVpnQueuePriority`

NewMsgVpnQueuePriorityWithDefaults instantiates a new MsgVpnQueuePriority object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsgVpnName

`func (o *MsgVpnQueuePriority) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnQueuePriority) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnQueuePriority) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnQueuePriority) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetPriority

`func (o *MsgVpnQueuePriority) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MsgVpnQueuePriority) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *MsgVpnQueuePriority) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *MsgVpnQueuePriority) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetQueueName

`func (o *MsgVpnQueuePriority) GetQueueName() string`

GetQueueName returns the QueueName field if non-nil, zero value otherwise.

### GetQueueNameOk

`func (o *MsgVpnQueuePriority) GetQueueNameOk() (*string, bool)`

GetQueueNameOk returns a tuple with the QueueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueName

`func (o *MsgVpnQueuePriority) SetQueueName(v string)`

SetQueueName sets QueueName field to given value.

### HasQueueName

`func (o *MsgVpnQueuePriority) HasQueueName() bool`

HasQueueName returns a boolean if a field has been set.

### GetSpooledByteCount

`func (o *MsgVpnQueuePriority) GetSpooledByteCount() int64`

GetSpooledByteCount returns the SpooledByteCount field if non-nil, zero value otherwise.

### GetSpooledByteCountOk

`func (o *MsgVpnQueuePriority) GetSpooledByteCountOk() (*int64, bool)`

GetSpooledByteCountOk returns a tuple with the SpooledByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpooledByteCount

`func (o *MsgVpnQueuePriority) SetSpooledByteCount(v int64)`

SetSpooledByteCount sets SpooledByteCount field to given value.

### HasSpooledByteCount

`func (o *MsgVpnQueuePriority) HasSpooledByteCount() bool`

HasSpooledByteCount returns a boolean if a field has been set.

### GetSpooledMsgCount

`func (o *MsgVpnQueuePriority) GetSpooledMsgCount() int64`

GetSpooledMsgCount returns the SpooledMsgCount field if non-nil, zero value otherwise.

### GetSpooledMsgCountOk

`func (o *MsgVpnQueuePriority) GetSpooledMsgCountOk() (*int64, bool)`

GetSpooledMsgCountOk returns a tuple with the SpooledMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpooledMsgCount

`func (o *MsgVpnQueuePriority) SetSpooledMsgCount(v int64)`

SetSpooledMsgCount sets SpooledMsgCount field to given value.

### HasSpooledMsgCount

`func (o *MsgVpnQueuePriority) HasSpooledMsgCount() bool`

HasSpooledMsgCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


