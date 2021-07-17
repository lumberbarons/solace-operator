# MsgVpnTransactionPublisherMsg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MsgId** | Pointer to **int64** | The identifier (ID) of the Message. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**Topic** | Pointer to **string** | The topic destination of the Message. | [optional] 
**Xid** | Pointer to **string** | The identifier (ID) of the Transaction. | [optional] 

## Methods

### NewMsgVpnTransactionPublisherMsg

`func NewMsgVpnTransactionPublisherMsg() *MsgVpnTransactionPublisherMsg`

NewMsgVpnTransactionPublisherMsg instantiates a new MsgVpnTransactionPublisherMsg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnTransactionPublisherMsgWithDefaults

`func NewMsgVpnTransactionPublisherMsgWithDefaults() *MsgVpnTransactionPublisherMsg`

NewMsgVpnTransactionPublisherMsgWithDefaults instantiates a new MsgVpnTransactionPublisherMsg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsgId

`func (o *MsgVpnTransactionPublisherMsg) GetMsgId() int64`

GetMsgId returns the MsgId field if non-nil, zero value otherwise.

### GetMsgIdOk

`func (o *MsgVpnTransactionPublisherMsg) GetMsgIdOk() (*int64, bool)`

GetMsgIdOk returns a tuple with the MsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgId

`func (o *MsgVpnTransactionPublisherMsg) SetMsgId(v int64)`

SetMsgId sets MsgId field to given value.

### HasMsgId

`func (o *MsgVpnTransactionPublisherMsg) HasMsgId() bool`

HasMsgId returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnTransactionPublisherMsg) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnTransactionPublisherMsg) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnTransactionPublisherMsg) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnTransactionPublisherMsg) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetTopic

`func (o *MsgVpnTransactionPublisherMsg) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *MsgVpnTransactionPublisherMsg) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *MsgVpnTransactionPublisherMsg) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *MsgVpnTransactionPublisherMsg) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetXid

`func (o *MsgVpnTransactionPublisherMsg) GetXid() string`

GetXid returns the Xid field if non-nil, zero value otherwise.

### GetXidOk

`func (o *MsgVpnTransactionPublisherMsg) GetXidOk() (*string, bool)`

GetXidOk returns a tuple with the Xid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXid

`func (o *MsgVpnTransactionPublisherMsg) SetXid(v string)`

SetXid sets Xid field to given value.

### HasXid

`func (o *MsgVpnTransactionPublisherMsg) HasXid() bool`

HasXid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


