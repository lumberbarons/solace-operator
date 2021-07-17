# MsgVpnTransactionConsumerMsg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndpointName** | Pointer to **string** | The name of the Queue or Topic Endpoint source. | [optional] 
**EndpointType** | Pointer to **string** | The type of endpoint source. The allowed values and their meaning are:  &lt;pre&gt; \&quot;queue\&quot; - The Message is from a Queue. \&quot;topic-endpoint\&quot; - The Message is from a Topic Endpoint. &lt;/pre&gt;  | [optional] 
**MsgId** | Pointer to **int64** | The identifier (ID) of the Message. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**ReplicationGroupMsgId** | Pointer to **string** | An ID that uniquely identifies this message within this replication group. Available since 2.21. | [optional] 
**Xid** | Pointer to **string** | The identifier (ID) of the Transaction. | [optional] 

## Methods

### NewMsgVpnTransactionConsumerMsg

`func NewMsgVpnTransactionConsumerMsg() *MsgVpnTransactionConsumerMsg`

NewMsgVpnTransactionConsumerMsg instantiates a new MsgVpnTransactionConsumerMsg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnTransactionConsumerMsgWithDefaults

`func NewMsgVpnTransactionConsumerMsgWithDefaults() *MsgVpnTransactionConsumerMsg`

NewMsgVpnTransactionConsumerMsgWithDefaults instantiates a new MsgVpnTransactionConsumerMsg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpointName

`func (o *MsgVpnTransactionConsumerMsg) GetEndpointName() string`

GetEndpointName returns the EndpointName field if non-nil, zero value otherwise.

### GetEndpointNameOk

`func (o *MsgVpnTransactionConsumerMsg) GetEndpointNameOk() (*string, bool)`

GetEndpointNameOk returns a tuple with the EndpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointName

`func (o *MsgVpnTransactionConsumerMsg) SetEndpointName(v string)`

SetEndpointName sets EndpointName field to given value.

### HasEndpointName

`func (o *MsgVpnTransactionConsumerMsg) HasEndpointName() bool`

HasEndpointName returns a boolean if a field has been set.

### GetEndpointType

`func (o *MsgVpnTransactionConsumerMsg) GetEndpointType() string`

GetEndpointType returns the EndpointType field if non-nil, zero value otherwise.

### GetEndpointTypeOk

`func (o *MsgVpnTransactionConsumerMsg) GetEndpointTypeOk() (*string, bool)`

GetEndpointTypeOk returns a tuple with the EndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointType

`func (o *MsgVpnTransactionConsumerMsg) SetEndpointType(v string)`

SetEndpointType sets EndpointType field to given value.

### HasEndpointType

`func (o *MsgVpnTransactionConsumerMsg) HasEndpointType() bool`

HasEndpointType returns a boolean if a field has been set.

### GetMsgId

`func (o *MsgVpnTransactionConsumerMsg) GetMsgId() int64`

GetMsgId returns the MsgId field if non-nil, zero value otherwise.

### GetMsgIdOk

`func (o *MsgVpnTransactionConsumerMsg) GetMsgIdOk() (*int64, bool)`

GetMsgIdOk returns a tuple with the MsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgId

`func (o *MsgVpnTransactionConsumerMsg) SetMsgId(v int64)`

SetMsgId sets MsgId field to given value.

### HasMsgId

`func (o *MsgVpnTransactionConsumerMsg) HasMsgId() bool`

HasMsgId returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnTransactionConsumerMsg) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnTransactionConsumerMsg) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnTransactionConsumerMsg) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnTransactionConsumerMsg) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetReplicationGroupMsgId

`func (o *MsgVpnTransactionConsumerMsg) GetReplicationGroupMsgId() string`

GetReplicationGroupMsgId returns the ReplicationGroupMsgId field if non-nil, zero value otherwise.

### GetReplicationGroupMsgIdOk

`func (o *MsgVpnTransactionConsumerMsg) GetReplicationGroupMsgIdOk() (*string, bool)`

GetReplicationGroupMsgIdOk returns a tuple with the ReplicationGroupMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationGroupMsgId

`func (o *MsgVpnTransactionConsumerMsg) SetReplicationGroupMsgId(v string)`

SetReplicationGroupMsgId sets ReplicationGroupMsgId field to given value.

### HasReplicationGroupMsgId

`func (o *MsgVpnTransactionConsumerMsg) HasReplicationGroupMsgId() bool`

HasReplicationGroupMsgId returns a boolean if a field has been set.

### GetXid

`func (o *MsgVpnTransactionConsumerMsg) GetXid() string`

GetXid returns the Xid field if non-nil, zero value otherwise.

### GetXidOk

`func (o *MsgVpnTransactionConsumerMsg) GetXidOk() (*string, bool)`

GetXidOk returns a tuple with the Xid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXid

`func (o *MsgVpnTransactionConsumerMsg) SetXid(v string)`

SetXid sets Xid field to given value.

### HasXid

`func (o *MsgVpnTransactionConsumerMsg) HasXid() bool`

HasXid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


