# MsgVpnTopicEndpointPriority

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**Priority** | Pointer to **int64** | The level of the Priority, from 9 (highest) to 0 (lowest). | [optional] 
**SpooledByteCount** | Pointer to **int64** | The amount of guaranteed messages at this Priority spooled by the Topic Endpoint, in bytes (B). | [optional] 
**SpooledMsgCount** | Pointer to **int64** | The number of guaranteed messages at this Priority spooled by the Topic Endpoint. | [optional] 
**TopicEndpointName** | Pointer to **string** | The name of the Topic Endpoint. | [optional] 

## Methods

### NewMsgVpnTopicEndpointPriority

`func NewMsgVpnTopicEndpointPriority() *MsgVpnTopicEndpointPriority`

NewMsgVpnTopicEndpointPriority instantiates a new MsgVpnTopicEndpointPriority object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnTopicEndpointPriorityWithDefaults

`func NewMsgVpnTopicEndpointPriorityWithDefaults() *MsgVpnTopicEndpointPriority`

NewMsgVpnTopicEndpointPriorityWithDefaults instantiates a new MsgVpnTopicEndpointPriority object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsgVpnName

`func (o *MsgVpnTopicEndpointPriority) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnTopicEndpointPriority) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnTopicEndpointPriority) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnTopicEndpointPriority) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetPriority

`func (o *MsgVpnTopicEndpointPriority) GetPriority() int64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MsgVpnTopicEndpointPriority) GetPriorityOk() (*int64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *MsgVpnTopicEndpointPriority) SetPriority(v int64)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *MsgVpnTopicEndpointPriority) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetSpooledByteCount

`func (o *MsgVpnTopicEndpointPriority) GetSpooledByteCount() int64`

GetSpooledByteCount returns the SpooledByteCount field if non-nil, zero value otherwise.

### GetSpooledByteCountOk

`func (o *MsgVpnTopicEndpointPriority) GetSpooledByteCountOk() (*int64, bool)`

GetSpooledByteCountOk returns a tuple with the SpooledByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpooledByteCount

`func (o *MsgVpnTopicEndpointPriority) SetSpooledByteCount(v int64)`

SetSpooledByteCount sets SpooledByteCount field to given value.

### HasSpooledByteCount

`func (o *MsgVpnTopicEndpointPriority) HasSpooledByteCount() bool`

HasSpooledByteCount returns a boolean if a field has been set.

### GetSpooledMsgCount

`func (o *MsgVpnTopicEndpointPriority) GetSpooledMsgCount() int64`

GetSpooledMsgCount returns the SpooledMsgCount field if non-nil, zero value otherwise.

### GetSpooledMsgCountOk

`func (o *MsgVpnTopicEndpointPriority) GetSpooledMsgCountOk() (*int64, bool)`

GetSpooledMsgCountOk returns a tuple with the SpooledMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpooledMsgCount

`func (o *MsgVpnTopicEndpointPriority) SetSpooledMsgCount(v int64)`

SetSpooledMsgCount sets SpooledMsgCount field to given value.

### HasSpooledMsgCount

`func (o *MsgVpnTopicEndpointPriority) HasSpooledMsgCount() bool`

HasSpooledMsgCount returns a boolean if a field has been set.

### GetTopicEndpointName

`func (o *MsgVpnTopicEndpointPriority) GetTopicEndpointName() string`

GetTopicEndpointName returns the TopicEndpointName field if non-nil, zero value otherwise.

### GetTopicEndpointNameOk

`func (o *MsgVpnTopicEndpointPriority) GetTopicEndpointNameOk() (*string, bool)`

GetTopicEndpointNameOk returns a tuple with the TopicEndpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicEndpointName

`func (o *MsgVpnTopicEndpointPriority) SetTopicEndpointName(v string)`

SetTopicEndpointName sets TopicEndpointName field to given value.

### HasTopicEndpointName

`func (o *MsgVpnTopicEndpointPriority) HasTopicEndpointName() bool`

HasTopicEndpointName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


