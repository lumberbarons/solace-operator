# MsgVpnTopicEndpointTxFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AckedMsgCount** | Pointer to **int64** | The number of guaranteed messages delivered and acknowledged by the consumer. | [optional] 
**ActivityState** | Pointer to **string** | The activity state of the Flow. The allowed values and their meaning are:  &lt;pre&gt; \&quot;active-browser\&quot; - The Flow is active as a browser. \&quot;active-consumer\&quot; - The Flow is active as a consumer. \&quot;inactive\&quot; - The Flow is inactive. &lt;/pre&gt;  | [optional] 
**BindTime** | Pointer to **int32** | The timestamp of when the Flow bound to the Topic Endpoint. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] 
**ClientName** | Pointer to **string** | The name of the Client. | [optional] 
**ConsumerRedeliveryRequestAllowed** | Pointer to **bool** | Indicates whether redelivery requests can be received as negative acknowledgements (NACKs) from the consumer. Applicable only to REST consumers. | [optional] 
**CutThroughAckedMsgCount** | Pointer to **int64** | The number of guaranteed messages that used cut-through delivery and are acknowledged by the consumer. | [optional] 
**DeliveryState** | Pointer to **string** | The delivery state of the Flow. The allowed values and their meaning are:  &lt;pre&gt; \&quot;closed\&quot; - The Flow is unbound. \&quot;opened\&quot; - The Flow is bound but inactive. \&quot;unbinding\&quot; - The Flow received an unbind request. \&quot;handshaking\&quot; - The Flow is handshaking to become active. \&quot;deliver-cut-through\&quot; - The Flow is streaming messages using direct+guaranteed delivery. \&quot;deliver-from-input-stream\&quot; - The Flow is streaming messages using guaranteed delivery. \&quot;deliver-from-memory\&quot; - The Flow throttled causing message delivery from memory (RAM). \&quot;deliver-from-spool\&quot; - The Flow stalled causing message delivery from spool (ADB or disk). &lt;/pre&gt;  | [optional] 
**FlowId** | Pointer to **int64** | The identifier (ID) of the Flow. | [optional] 
**HighestAckPendingMsgId** | Pointer to **int64** | The highest identifier (ID) of message transmitted and waiting for acknowledgement. | [optional] 
**LastAckedMsgId** | Pointer to **int64** | The identifier (ID) of the last message transmitted and acknowledged by the consumer. | [optional] 
**LowestAckPendingMsgId** | Pointer to **int64** | The lowest identifier (ID) of message transmitted and waiting for acknowledgement. | [optional] 
**MaxUnackedMsgsExceededMsgCount** | Pointer to **int64** | The number of guaranteed messages that exceeded the maximum number of delivered unacknowledged messages. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**NoLocalDelivery** | Pointer to **bool** | Indicates whether not to deliver messages to a consumer that published them. | [optional] 
**RedeliveredMsgCount** | Pointer to **int64** | The number of guaranteed messages that were redelivered. | [optional] 
**RedeliveryRequestCount** | Pointer to **int64** | The number of consumer requests via negative acknowledgements (NACKs) to redeliver guaranteed messages. | [optional] 
**SessionName** | Pointer to **string** | The name of the Transacted Session for the Flow. | [optional] 
**StoreAndForwardAckedMsgCount** | Pointer to **int64** | The number of guaranteed messages that used store and forward delivery and are acknowledged by the consumer. | [optional] 
**TopicEndpointName** | Pointer to **string** | The name of the Topic Endpoint. | [optional] 
**TransportRetransmitMsgCount** | Pointer to **int64** | The number of guaranteed messages that were retransmitted at the transport layer as part of a single delivery attempt. Available since 2.18. | [optional] 
**UnackedMsgCount** | Pointer to **int64** | The number of guaranteed messages delivered but not yet acknowledged by the consumer. | [optional] 
**UsedWindowSize** | Pointer to **int64** | The number of guaranteed messages using the available window size. | [optional] 
**WindowClosedCount** | Pointer to **int64** | The number of times the window for guaranteed messages was filled and closed before an acknowledgement was received. | [optional] 
**WindowSize** | Pointer to **int64** | The number of outstanding guaranteed messages that can be transmitted over the Flow before an acknowledgement is received. | [optional] 

## Methods

### NewMsgVpnTopicEndpointTxFlow

`func NewMsgVpnTopicEndpointTxFlow() *MsgVpnTopicEndpointTxFlow`

NewMsgVpnTopicEndpointTxFlow instantiates a new MsgVpnTopicEndpointTxFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnTopicEndpointTxFlowWithDefaults

`func NewMsgVpnTopicEndpointTxFlowWithDefaults() *MsgVpnTopicEndpointTxFlow`

NewMsgVpnTopicEndpointTxFlowWithDefaults instantiates a new MsgVpnTopicEndpointTxFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAckedMsgCount

`func (o *MsgVpnTopicEndpointTxFlow) GetAckedMsgCount() int64`

GetAckedMsgCount returns the AckedMsgCount field if non-nil, zero value otherwise.

### GetAckedMsgCountOk

`func (o *MsgVpnTopicEndpointTxFlow) GetAckedMsgCountOk() (*int64, bool)`

GetAckedMsgCountOk returns a tuple with the AckedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckedMsgCount

`func (o *MsgVpnTopicEndpointTxFlow) SetAckedMsgCount(v int64)`

SetAckedMsgCount sets AckedMsgCount field to given value.

### HasAckedMsgCount

`func (o *MsgVpnTopicEndpointTxFlow) HasAckedMsgCount() bool`

HasAckedMsgCount returns a boolean if a field has been set.

### GetActivityState

`func (o *MsgVpnTopicEndpointTxFlow) GetActivityState() string`

GetActivityState returns the ActivityState field if non-nil, zero value otherwise.

### GetActivityStateOk

`func (o *MsgVpnTopicEndpointTxFlow) GetActivityStateOk() (*string, bool)`

GetActivityStateOk returns a tuple with the ActivityState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityState

`func (o *MsgVpnTopicEndpointTxFlow) SetActivityState(v string)`

SetActivityState sets ActivityState field to given value.

### HasActivityState

`func (o *MsgVpnTopicEndpointTxFlow) HasActivityState() bool`

HasActivityState returns a boolean if a field has been set.

### GetBindTime

`func (o *MsgVpnTopicEndpointTxFlow) GetBindTime() int32`

GetBindTime returns the BindTime field if non-nil, zero value otherwise.

### GetBindTimeOk

`func (o *MsgVpnTopicEndpointTxFlow) GetBindTimeOk() (*int32, bool)`

GetBindTimeOk returns a tuple with the BindTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindTime

`func (o *MsgVpnTopicEndpointTxFlow) SetBindTime(v int32)`

SetBindTime sets BindTime field to given value.

### HasBindTime

`func (o *MsgVpnTopicEndpointTxFlow) HasBindTime() bool`

HasBindTime returns a boolean if a field has been set.

### GetClientName

`func (o *MsgVpnTopicEndpointTxFlow) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *MsgVpnTopicEndpointTxFlow) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *MsgVpnTopicEndpointTxFlow) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *MsgVpnTopicEndpointTxFlow) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetConsumerRedeliveryRequestAllowed

`func (o *MsgVpnTopicEndpointTxFlow) GetConsumerRedeliveryRequestAllowed() bool`

GetConsumerRedeliveryRequestAllowed returns the ConsumerRedeliveryRequestAllowed field if non-nil, zero value otherwise.

### GetConsumerRedeliveryRequestAllowedOk

`func (o *MsgVpnTopicEndpointTxFlow) GetConsumerRedeliveryRequestAllowedOk() (*bool, bool)`

GetConsumerRedeliveryRequestAllowedOk returns a tuple with the ConsumerRedeliveryRequestAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerRedeliveryRequestAllowed

`func (o *MsgVpnTopicEndpointTxFlow) SetConsumerRedeliveryRequestAllowed(v bool)`

SetConsumerRedeliveryRequestAllowed sets ConsumerRedeliveryRequestAllowed field to given value.

### HasConsumerRedeliveryRequestAllowed

`func (o *MsgVpnTopicEndpointTxFlow) HasConsumerRedeliveryRequestAllowed() bool`

HasConsumerRedeliveryRequestAllowed returns a boolean if a field has been set.

### GetCutThroughAckedMsgCount

`func (o *MsgVpnTopicEndpointTxFlow) GetCutThroughAckedMsgCount() int64`

GetCutThroughAckedMsgCount returns the CutThroughAckedMsgCount field if non-nil, zero value otherwise.

### GetCutThroughAckedMsgCountOk

`func (o *MsgVpnTopicEndpointTxFlow) GetCutThroughAckedMsgCountOk() (*int64, bool)`

GetCutThroughAckedMsgCountOk returns a tuple with the CutThroughAckedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCutThroughAckedMsgCount

`func (o *MsgVpnTopicEndpointTxFlow) SetCutThroughAckedMsgCount(v int64)`

SetCutThroughAckedMsgCount sets CutThroughAckedMsgCount field to given value.

### HasCutThroughAckedMsgCount

`func (o *MsgVpnTopicEndpointTxFlow) HasCutThroughAckedMsgCount() bool`

HasCutThroughAckedMsgCount returns a boolean if a field has been set.

### GetDeliveryState

`func (o *MsgVpnTopicEndpointTxFlow) GetDeliveryState() string`

GetDeliveryState returns the DeliveryState field if non-nil, zero value otherwise.

### GetDeliveryStateOk

`func (o *MsgVpnTopicEndpointTxFlow) GetDeliveryStateOk() (*string, bool)`

GetDeliveryStateOk returns a tuple with the DeliveryState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryState

`func (o *MsgVpnTopicEndpointTxFlow) SetDeliveryState(v string)`

SetDeliveryState sets DeliveryState field to given value.

### HasDeliveryState

`func (o *MsgVpnTopicEndpointTxFlow) HasDeliveryState() bool`

HasDeliveryState returns a boolean if a field has been set.

### GetFlowId

`func (o *MsgVpnTopicEndpointTxFlow) GetFlowId() int64`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *MsgVpnTopicEndpointTxFlow) GetFlowIdOk() (*int64, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *MsgVpnTopicEndpointTxFlow) SetFlowId(v int64)`

SetFlowId sets FlowId field to given value.

### HasFlowId

`func (o *MsgVpnTopicEndpointTxFlow) HasFlowId() bool`

HasFlowId returns a boolean if a field has been set.

### GetHighestAckPendingMsgId

`func (o *MsgVpnTopicEndpointTxFlow) GetHighestAckPendingMsgId() int64`

GetHighestAckPendingMsgId returns the HighestAckPendingMsgId field if non-nil, zero value otherwise.

### GetHighestAckPendingMsgIdOk

`func (o *MsgVpnTopicEndpointTxFlow) GetHighestAckPendingMsgIdOk() (*int64, bool)`

GetHighestAckPendingMsgIdOk returns a tuple with the HighestAckPendingMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestAckPendingMsgId

`func (o *MsgVpnTopicEndpointTxFlow) SetHighestAckPendingMsgId(v int64)`

SetHighestAckPendingMsgId sets HighestAckPendingMsgId field to given value.

### HasHighestAckPendingMsgId

`func (o *MsgVpnTopicEndpointTxFlow) HasHighestAckPendingMsgId() bool`

HasHighestAckPendingMsgId returns a boolean if a field has been set.

### GetLastAckedMsgId

`func (o *MsgVpnTopicEndpointTxFlow) GetLastAckedMsgId() int64`

GetLastAckedMsgId returns the LastAckedMsgId field if non-nil, zero value otherwise.

### GetLastAckedMsgIdOk

`func (o *MsgVpnTopicEndpointTxFlow) GetLastAckedMsgIdOk() (*int64, bool)`

GetLastAckedMsgIdOk returns a tuple with the LastAckedMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAckedMsgId

`func (o *MsgVpnTopicEndpointTxFlow) SetLastAckedMsgId(v int64)`

SetLastAckedMsgId sets LastAckedMsgId field to given value.

### HasLastAckedMsgId

`func (o *MsgVpnTopicEndpointTxFlow) HasLastAckedMsgId() bool`

HasLastAckedMsgId returns a boolean if a field has been set.

### GetLowestAckPendingMsgId

`func (o *MsgVpnTopicEndpointTxFlow) GetLowestAckPendingMsgId() int64`

GetLowestAckPendingMsgId returns the LowestAckPendingMsgId field if non-nil, zero value otherwise.

### GetLowestAckPendingMsgIdOk

`func (o *MsgVpnTopicEndpointTxFlow) GetLowestAckPendingMsgIdOk() (*int64, bool)`

GetLowestAckPendingMsgIdOk returns a tuple with the LowestAckPendingMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowestAckPendingMsgId

`func (o *MsgVpnTopicEndpointTxFlow) SetLowestAckPendingMsgId(v int64)`

SetLowestAckPendingMsgId sets LowestAckPendingMsgId field to given value.

### HasLowestAckPendingMsgId

`func (o *MsgVpnTopicEndpointTxFlow) HasLowestAckPendingMsgId() bool`

HasLowestAckPendingMsgId returns a boolean if a field has been set.

### GetMaxUnackedMsgsExceededMsgCount

`func (o *MsgVpnTopicEndpointTxFlow) GetMaxUnackedMsgsExceededMsgCount() int64`

GetMaxUnackedMsgsExceededMsgCount returns the MaxUnackedMsgsExceededMsgCount field if non-nil, zero value otherwise.

### GetMaxUnackedMsgsExceededMsgCountOk

`func (o *MsgVpnTopicEndpointTxFlow) GetMaxUnackedMsgsExceededMsgCountOk() (*int64, bool)`

GetMaxUnackedMsgsExceededMsgCountOk returns a tuple with the MaxUnackedMsgsExceededMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUnackedMsgsExceededMsgCount

`func (o *MsgVpnTopicEndpointTxFlow) SetMaxUnackedMsgsExceededMsgCount(v int64)`

SetMaxUnackedMsgsExceededMsgCount sets MaxUnackedMsgsExceededMsgCount field to given value.

### HasMaxUnackedMsgsExceededMsgCount

`func (o *MsgVpnTopicEndpointTxFlow) HasMaxUnackedMsgsExceededMsgCount() bool`

HasMaxUnackedMsgsExceededMsgCount returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnTopicEndpointTxFlow) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnTopicEndpointTxFlow) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnTopicEndpointTxFlow) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnTopicEndpointTxFlow) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetNoLocalDelivery

`func (o *MsgVpnTopicEndpointTxFlow) GetNoLocalDelivery() bool`

GetNoLocalDelivery returns the NoLocalDelivery field if non-nil, zero value otherwise.

### GetNoLocalDeliveryOk

`func (o *MsgVpnTopicEndpointTxFlow) GetNoLocalDeliveryOk() (*bool, bool)`

GetNoLocalDeliveryOk returns a tuple with the NoLocalDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoLocalDelivery

`func (o *MsgVpnTopicEndpointTxFlow) SetNoLocalDelivery(v bool)`

SetNoLocalDelivery sets NoLocalDelivery field to given value.

### HasNoLocalDelivery

`func (o *MsgVpnTopicEndpointTxFlow) HasNoLocalDelivery() bool`

HasNoLocalDelivery returns a boolean if a field has been set.

### GetRedeliveredMsgCount

`func (o *MsgVpnTopicEndpointTxFlow) GetRedeliveredMsgCount() int64`

GetRedeliveredMsgCount returns the RedeliveredMsgCount field if non-nil, zero value otherwise.

### GetRedeliveredMsgCountOk

`func (o *MsgVpnTopicEndpointTxFlow) GetRedeliveredMsgCountOk() (*int64, bool)`

GetRedeliveredMsgCountOk returns a tuple with the RedeliveredMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeliveredMsgCount

`func (o *MsgVpnTopicEndpointTxFlow) SetRedeliveredMsgCount(v int64)`

SetRedeliveredMsgCount sets RedeliveredMsgCount field to given value.

### HasRedeliveredMsgCount

`func (o *MsgVpnTopicEndpointTxFlow) HasRedeliveredMsgCount() bool`

HasRedeliveredMsgCount returns a boolean if a field has been set.

### GetRedeliveryRequestCount

`func (o *MsgVpnTopicEndpointTxFlow) GetRedeliveryRequestCount() int64`

GetRedeliveryRequestCount returns the RedeliveryRequestCount field if non-nil, zero value otherwise.

### GetRedeliveryRequestCountOk

`func (o *MsgVpnTopicEndpointTxFlow) GetRedeliveryRequestCountOk() (*int64, bool)`

GetRedeliveryRequestCountOk returns a tuple with the RedeliveryRequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeliveryRequestCount

`func (o *MsgVpnTopicEndpointTxFlow) SetRedeliveryRequestCount(v int64)`

SetRedeliveryRequestCount sets RedeliveryRequestCount field to given value.

### HasRedeliveryRequestCount

`func (o *MsgVpnTopicEndpointTxFlow) HasRedeliveryRequestCount() bool`

HasRedeliveryRequestCount returns a boolean if a field has been set.

### GetSessionName

`func (o *MsgVpnTopicEndpointTxFlow) GetSessionName() string`

GetSessionName returns the SessionName field if non-nil, zero value otherwise.

### GetSessionNameOk

`func (o *MsgVpnTopicEndpointTxFlow) GetSessionNameOk() (*string, bool)`

GetSessionNameOk returns a tuple with the SessionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionName

`func (o *MsgVpnTopicEndpointTxFlow) SetSessionName(v string)`

SetSessionName sets SessionName field to given value.

### HasSessionName

`func (o *MsgVpnTopicEndpointTxFlow) HasSessionName() bool`

HasSessionName returns a boolean if a field has been set.

### GetStoreAndForwardAckedMsgCount

`func (o *MsgVpnTopicEndpointTxFlow) GetStoreAndForwardAckedMsgCount() int64`

GetStoreAndForwardAckedMsgCount returns the StoreAndForwardAckedMsgCount field if non-nil, zero value otherwise.

### GetStoreAndForwardAckedMsgCountOk

`func (o *MsgVpnTopicEndpointTxFlow) GetStoreAndForwardAckedMsgCountOk() (*int64, bool)`

GetStoreAndForwardAckedMsgCountOk returns a tuple with the StoreAndForwardAckedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreAndForwardAckedMsgCount

`func (o *MsgVpnTopicEndpointTxFlow) SetStoreAndForwardAckedMsgCount(v int64)`

SetStoreAndForwardAckedMsgCount sets StoreAndForwardAckedMsgCount field to given value.

### HasStoreAndForwardAckedMsgCount

`func (o *MsgVpnTopicEndpointTxFlow) HasStoreAndForwardAckedMsgCount() bool`

HasStoreAndForwardAckedMsgCount returns a boolean if a field has been set.

### GetTopicEndpointName

`func (o *MsgVpnTopicEndpointTxFlow) GetTopicEndpointName() string`

GetTopicEndpointName returns the TopicEndpointName field if non-nil, zero value otherwise.

### GetTopicEndpointNameOk

`func (o *MsgVpnTopicEndpointTxFlow) GetTopicEndpointNameOk() (*string, bool)`

GetTopicEndpointNameOk returns a tuple with the TopicEndpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicEndpointName

`func (o *MsgVpnTopicEndpointTxFlow) SetTopicEndpointName(v string)`

SetTopicEndpointName sets TopicEndpointName field to given value.

### HasTopicEndpointName

`func (o *MsgVpnTopicEndpointTxFlow) HasTopicEndpointName() bool`

HasTopicEndpointName returns a boolean if a field has been set.

### GetTransportRetransmitMsgCount

`func (o *MsgVpnTopicEndpointTxFlow) GetTransportRetransmitMsgCount() int64`

GetTransportRetransmitMsgCount returns the TransportRetransmitMsgCount field if non-nil, zero value otherwise.

### GetTransportRetransmitMsgCountOk

`func (o *MsgVpnTopicEndpointTxFlow) GetTransportRetransmitMsgCountOk() (*int64, bool)`

GetTransportRetransmitMsgCountOk returns a tuple with the TransportRetransmitMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportRetransmitMsgCount

`func (o *MsgVpnTopicEndpointTxFlow) SetTransportRetransmitMsgCount(v int64)`

SetTransportRetransmitMsgCount sets TransportRetransmitMsgCount field to given value.

### HasTransportRetransmitMsgCount

`func (o *MsgVpnTopicEndpointTxFlow) HasTransportRetransmitMsgCount() bool`

HasTransportRetransmitMsgCount returns a boolean if a field has been set.

### GetUnackedMsgCount

`func (o *MsgVpnTopicEndpointTxFlow) GetUnackedMsgCount() int64`

GetUnackedMsgCount returns the UnackedMsgCount field if non-nil, zero value otherwise.

### GetUnackedMsgCountOk

`func (o *MsgVpnTopicEndpointTxFlow) GetUnackedMsgCountOk() (*int64, bool)`

GetUnackedMsgCountOk returns a tuple with the UnackedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnackedMsgCount

`func (o *MsgVpnTopicEndpointTxFlow) SetUnackedMsgCount(v int64)`

SetUnackedMsgCount sets UnackedMsgCount field to given value.

### HasUnackedMsgCount

`func (o *MsgVpnTopicEndpointTxFlow) HasUnackedMsgCount() bool`

HasUnackedMsgCount returns a boolean if a field has been set.

### GetUsedWindowSize

`func (o *MsgVpnTopicEndpointTxFlow) GetUsedWindowSize() int64`

GetUsedWindowSize returns the UsedWindowSize field if non-nil, zero value otherwise.

### GetUsedWindowSizeOk

`func (o *MsgVpnTopicEndpointTxFlow) GetUsedWindowSizeOk() (*int64, bool)`

GetUsedWindowSizeOk returns a tuple with the UsedWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedWindowSize

`func (o *MsgVpnTopicEndpointTxFlow) SetUsedWindowSize(v int64)`

SetUsedWindowSize sets UsedWindowSize field to given value.

### HasUsedWindowSize

`func (o *MsgVpnTopicEndpointTxFlow) HasUsedWindowSize() bool`

HasUsedWindowSize returns a boolean if a field has been set.

### GetWindowClosedCount

`func (o *MsgVpnTopicEndpointTxFlow) GetWindowClosedCount() int64`

GetWindowClosedCount returns the WindowClosedCount field if non-nil, zero value otherwise.

### GetWindowClosedCountOk

`func (o *MsgVpnTopicEndpointTxFlow) GetWindowClosedCountOk() (*int64, bool)`

GetWindowClosedCountOk returns a tuple with the WindowClosedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowClosedCount

`func (o *MsgVpnTopicEndpointTxFlow) SetWindowClosedCount(v int64)`

SetWindowClosedCount sets WindowClosedCount field to given value.

### HasWindowClosedCount

`func (o *MsgVpnTopicEndpointTxFlow) HasWindowClosedCount() bool`

HasWindowClosedCount returns a boolean if a field has been set.

### GetWindowSize

`func (o *MsgVpnTopicEndpointTxFlow) GetWindowSize() int64`

GetWindowSize returns the WindowSize field if non-nil, zero value otherwise.

### GetWindowSizeOk

`func (o *MsgVpnTopicEndpointTxFlow) GetWindowSizeOk() (*int64, bool)`

GetWindowSizeOk returns a tuple with the WindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSize

`func (o *MsgVpnTopicEndpointTxFlow) SetWindowSize(v int64)`

SetWindowSize sets WindowSize field to given value.

### HasWindowSize

`func (o *MsgVpnTopicEndpointTxFlow) HasWindowSize() bool`

HasWindowSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


