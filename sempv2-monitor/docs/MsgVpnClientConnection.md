# MsgVpnClientConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientAddress** | Pointer to **string** | The IP address and TCP port on the Client side of the Client Connection. | [optional] 
**ClientName** | Pointer to **string** | The name of the Client. | [optional] 
**Compression** | Pointer to **bool** | Indicates whether compression is enabled for the Client Connection. | [optional] 
**Encryption** | Pointer to **bool** | Indicates whether encryption (TLS) is enabled for the Client Connection. | [optional] 
**FastRetransmitCount** | Pointer to **int32** | The number of TCP fast retransmits due to duplicate acknowledgments (ACKs). See RFC 5681 for further details. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**RxQueueByteCount** | Pointer to **int32** | The number of bytes currently in the receive queue for the Client Connection. | [optional] 
**SegmentReceivedOutOfOrderCount** | Pointer to **int32** | The number of TCP segments received from the Client Connection out of order. | [optional] 
**SmoothedRoundTripTime** | Pointer to **int64** | The TCP smoothed round-trip time (SRTT) for the Client Connection, in nanoseconds. See RFC 2988 for further details. | [optional] 
**TcpState** | Pointer to **string** | The TCP state of the Client Connection. When fully operational, should be: established. See RFC 793 for further details. The allowed values and their meaning are:  &lt;pre&gt; \&quot;closed\&quot; - No connection state at all. \&quot;listen\&quot; - Waiting for a connection request from any remote TCP and port. \&quot;syn-sent\&quot; - Waiting for a matching connection request after having sent a connection request. \&quot;syn-received\&quot; - Waiting for a confirming connection request acknowledgment after having both received and sent a connection request. \&quot;established\&quot; - An open connection, data received can be delivered to the user. \&quot;close-wait\&quot; - Waiting for a connection termination request from the local user. \&quot;fin-wait-1\&quot; - Waiting for a connection termination request from the remote TCP, or an acknowledgment of the connection termination request previously sent. \&quot;closing\&quot; - Waiting for a connection termination request acknowledgment from the remote TCP. \&quot;last-ack\&quot; - Waiting for an acknowledgment of the connection termination request previously sent to the remote TCP. \&quot;fin-wait-2\&quot; - Waiting for a connection termination request from the remote TCP. \&quot;time-wait\&quot; - Waiting for enough time to pass to be sure the remote TCP received the acknowledgment of its connection termination request. &lt;/pre&gt;  | [optional] 
**TimedRetransmitCount** | Pointer to **int32** | The number of TCP segments retransmitted due to timeout awaiting an acknowledgement (ACK). See RFC 793 for further details. | [optional] 
**TxQueueByteCount** | Pointer to **int32** | The number of bytes currently in the transmit queue for the Client Connection. | [optional] 
**Uptime** | Pointer to **int64** | The amount of time in seconds since the Client Connection was established. | [optional] 

## Methods

### NewMsgVpnClientConnection

`func NewMsgVpnClientConnection() *MsgVpnClientConnection`

NewMsgVpnClientConnection instantiates a new MsgVpnClientConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnClientConnectionWithDefaults

`func NewMsgVpnClientConnectionWithDefaults() *MsgVpnClientConnection`

NewMsgVpnClientConnectionWithDefaults instantiates a new MsgVpnClientConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientAddress

`func (o *MsgVpnClientConnection) GetClientAddress() string`

GetClientAddress returns the ClientAddress field if non-nil, zero value otherwise.

### GetClientAddressOk

`func (o *MsgVpnClientConnection) GetClientAddressOk() (*string, bool)`

GetClientAddressOk returns a tuple with the ClientAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAddress

`func (o *MsgVpnClientConnection) SetClientAddress(v string)`

SetClientAddress sets ClientAddress field to given value.

### HasClientAddress

`func (o *MsgVpnClientConnection) HasClientAddress() bool`

HasClientAddress returns a boolean if a field has been set.

### GetClientName

`func (o *MsgVpnClientConnection) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *MsgVpnClientConnection) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *MsgVpnClientConnection) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *MsgVpnClientConnection) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetCompression

`func (o *MsgVpnClientConnection) GetCompression() bool`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *MsgVpnClientConnection) GetCompressionOk() (*bool, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *MsgVpnClientConnection) SetCompression(v bool)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *MsgVpnClientConnection) HasCompression() bool`

HasCompression returns a boolean if a field has been set.

### GetEncryption

`func (o *MsgVpnClientConnection) GetEncryption() bool`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *MsgVpnClientConnection) GetEncryptionOk() (*bool, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *MsgVpnClientConnection) SetEncryption(v bool)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *MsgVpnClientConnection) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetFastRetransmitCount

`func (o *MsgVpnClientConnection) GetFastRetransmitCount() int32`

GetFastRetransmitCount returns the FastRetransmitCount field if non-nil, zero value otherwise.

### GetFastRetransmitCountOk

`func (o *MsgVpnClientConnection) GetFastRetransmitCountOk() (*int32, bool)`

GetFastRetransmitCountOk returns a tuple with the FastRetransmitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastRetransmitCount

`func (o *MsgVpnClientConnection) SetFastRetransmitCount(v int32)`

SetFastRetransmitCount sets FastRetransmitCount field to given value.

### HasFastRetransmitCount

`func (o *MsgVpnClientConnection) HasFastRetransmitCount() bool`

HasFastRetransmitCount returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnClientConnection) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnClientConnection) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnClientConnection) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnClientConnection) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetRxQueueByteCount

`func (o *MsgVpnClientConnection) GetRxQueueByteCount() int32`

GetRxQueueByteCount returns the RxQueueByteCount field if non-nil, zero value otherwise.

### GetRxQueueByteCountOk

`func (o *MsgVpnClientConnection) GetRxQueueByteCountOk() (*int32, bool)`

GetRxQueueByteCountOk returns a tuple with the RxQueueByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxQueueByteCount

`func (o *MsgVpnClientConnection) SetRxQueueByteCount(v int32)`

SetRxQueueByteCount sets RxQueueByteCount field to given value.

### HasRxQueueByteCount

`func (o *MsgVpnClientConnection) HasRxQueueByteCount() bool`

HasRxQueueByteCount returns a boolean if a field has been set.

### GetSegmentReceivedOutOfOrderCount

`func (o *MsgVpnClientConnection) GetSegmentReceivedOutOfOrderCount() int32`

GetSegmentReceivedOutOfOrderCount returns the SegmentReceivedOutOfOrderCount field if non-nil, zero value otherwise.

### GetSegmentReceivedOutOfOrderCountOk

`func (o *MsgVpnClientConnection) GetSegmentReceivedOutOfOrderCountOk() (*int32, bool)`

GetSegmentReceivedOutOfOrderCountOk returns a tuple with the SegmentReceivedOutOfOrderCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentReceivedOutOfOrderCount

`func (o *MsgVpnClientConnection) SetSegmentReceivedOutOfOrderCount(v int32)`

SetSegmentReceivedOutOfOrderCount sets SegmentReceivedOutOfOrderCount field to given value.

### HasSegmentReceivedOutOfOrderCount

`func (o *MsgVpnClientConnection) HasSegmentReceivedOutOfOrderCount() bool`

HasSegmentReceivedOutOfOrderCount returns a boolean if a field has been set.

### GetSmoothedRoundTripTime

`func (o *MsgVpnClientConnection) GetSmoothedRoundTripTime() int64`

GetSmoothedRoundTripTime returns the SmoothedRoundTripTime field if non-nil, zero value otherwise.

### GetSmoothedRoundTripTimeOk

`func (o *MsgVpnClientConnection) GetSmoothedRoundTripTimeOk() (*int64, bool)`

GetSmoothedRoundTripTimeOk returns a tuple with the SmoothedRoundTripTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmoothedRoundTripTime

`func (o *MsgVpnClientConnection) SetSmoothedRoundTripTime(v int64)`

SetSmoothedRoundTripTime sets SmoothedRoundTripTime field to given value.

### HasSmoothedRoundTripTime

`func (o *MsgVpnClientConnection) HasSmoothedRoundTripTime() bool`

HasSmoothedRoundTripTime returns a boolean if a field has been set.

### GetTcpState

`func (o *MsgVpnClientConnection) GetTcpState() string`

GetTcpState returns the TcpState field if non-nil, zero value otherwise.

### GetTcpStateOk

`func (o *MsgVpnClientConnection) GetTcpStateOk() (*string, bool)`

GetTcpStateOk returns a tuple with the TcpState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpState

`func (o *MsgVpnClientConnection) SetTcpState(v string)`

SetTcpState sets TcpState field to given value.

### HasTcpState

`func (o *MsgVpnClientConnection) HasTcpState() bool`

HasTcpState returns a boolean if a field has been set.

### GetTimedRetransmitCount

`func (o *MsgVpnClientConnection) GetTimedRetransmitCount() int32`

GetTimedRetransmitCount returns the TimedRetransmitCount field if non-nil, zero value otherwise.

### GetTimedRetransmitCountOk

`func (o *MsgVpnClientConnection) GetTimedRetransmitCountOk() (*int32, bool)`

GetTimedRetransmitCountOk returns a tuple with the TimedRetransmitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimedRetransmitCount

`func (o *MsgVpnClientConnection) SetTimedRetransmitCount(v int32)`

SetTimedRetransmitCount sets TimedRetransmitCount field to given value.

### HasTimedRetransmitCount

`func (o *MsgVpnClientConnection) HasTimedRetransmitCount() bool`

HasTimedRetransmitCount returns a boolean if a field has been set.

### GetTxQueueByteCount

`func (o *MsgVpnClientConnection) GetTxQueueByteCount() int32`

GetTxQueueByteCount returns the TxQueueByteCount field if non-nil, zero value otherwise.

### GetTxQueueByteCountOk

`func (o *MsgVpnClientConnection) GetTxQueueByteCountOk() (*int32, bool)`

GetTxQueueByteCountOk returns a tuple with the TxQueueByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxQueueByteCount

`func (o *MsgVpnClientConnection) SetTxQueueByteCount(v int32)`

SetTxQueueByteCount sets TxQueueByteCount field to given value.

### HasTxQueueByteCount

`func (o *MsgVpnClientConnection) HasTxQueueByteCount() bool`

HasTxQueueByteCount returns a boolean if a field has been set.

### GetUptime

`func (o *MsgVpnClientConnection) GetUptime() int64`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *MsgVpnClientConnection) GetUptimeOk() (*int64, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *MsgVpnClientConnection) SetUptime(v int64)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *MsgVpnClientConnection) HasUptime() bool`

HasUptime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


