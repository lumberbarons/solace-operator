# MsgVpnQueueTxFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AckedMsgCount** | Pointer to **int64** | The number of guaranteed messages delivered and acknowledged by the consumer. | [optional] 
**ActivationTime** | Pointer to **int32** | The timestamp of when the bound Flow became active. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] 
**ActivityState** | Pointer to **string** | The activity state of the Flow. The allowed values and their meaning are:  &lt;pre&gt; \&quot;active-browser\&quot; - The Flow is active as a browser. \&quot;active-consumer\&quot; - The Flow is active as a consumer. \&quot;inactive\&quot; - The Flow is inactive. &lt;/pre&gt;  | [optional] 
**ActivityUpdateState** | Pointer to **string** | The state of updating the consumer with the Flow activity. The allowed values and their meaning are:  &lt;pre&gt; \&quot;in-progress\&quot; - The Flow is in the process of updating the client with its activity state. \&quot;synchronized\&quot; - The Flow has updated the client with its activity state. \&quot;not-requested\&quot; - The Flow has not been requested by the client to provide activity updates. &lt;/pre&gt;  | [optional] 
**BindTime** | Pointer to **int32** | The timestamp of when the Flow bound to the Queue. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] 
**ClientName** | Pointer to **string** | The name of the Client. | [optional] 
**ConsumerRedeliveryRequestAllowed** | Pointer to **bool** | Indicates whether redelivery requests can be received as negative acknowledgements (NACKs) from the consumer. Applicable only to REST consumers. | [optional] 
**CutThroughAckedMsgCount** | Pointer to **int64** | The number of guaranteed messages that used cut-through delivery and are acknowledged by the consumer. | [optional] 
**DeliveryState** | Pointer to **string** | The delivery state of the Flow. The allowed values and their meaning are:  &lt;pre&gt; \&quot;closed\&quot; - The Flow is unbound. \&quot;opened\&quot; - The Flow is bound but inactive. \&quot;unbinding\&quot; - The Flow received an unbind request. \&quot;handshaking\&quot; - The Flow is handshaking to become active. \&quot;deliver-cut-through\&quot; - The Flow is streaming messages using direct+guaranteed delivery. \&quot;deliver-from-input-stream\&quot; - The Flow is streaming messages using guaranteed delivery. \&quot;deliver-from-memory\&quot; - The Flow throttled causing message delivery from memory (RAM). \&quot;deliver-from-spool\&quot; - The Flow stalled causing message delivery from spool (ADB or disk). &lt;/pre&gt;  | [optional] 
**FlowId** | Pointer to **int64** | The identifier (ID) of the Flow. | [optional] 
**HighestAckPendingMsgId** | Pointer to **int64** | The highest identifier (ID) of message transmitted and waiting for acknowledgement. | [optional] 
**LastAckedMsgId** | Pointer to **int64** | The identifier (ID) of the last message transmitted and acknowledged by the consumer. | [optional] 
**LastSelectorExaminedMsgId** | Pointer to **int64** | The identifier (ID) of the last message examined by the Flow selector. | [optional] 
**LowestAckPendingMsgId** | Pointer to **int64** | The lowest identifier (ID) of message transmitted and waiting for acknowledgement. | [optional] 
**MaxUnackedMsgsExceededMsgCount** | Pointer to **int64** | The number of guaranteed messages that exceeded the maximum number of delivered unacknowledged messages. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**NoLocalDelivery** | Pointer to **bool** | Indicates whether not to deliver messages to a consumer that published them. | [optional] 
**QueueName** | Pointer to **string** | The name of the Queue. | [optional] 
**RedeliveredMsgCount** | Pointer to **int64** | The number of guaranteed messages that were redelivered. | [optional] 
**RedeliveryRequestCount** | Pointer to **int64** | The number of consumer requests via negative acknowledgements (NACKs) to redeliver guaranteed messages. | [optional] 
**Selector** | Pointer to **string** | The value of the Flow selector. | [optional] 
**SelectorExaminedMsgCount** | Pointer to **int64** | The number of guaranteed messages examined by the Flow selector. | [optional] 
**SelectorMatchedMsgCount** | Pointer to **int64** | The number of guaranteed messages for which the Flow selector matched. | [optional] 
**SelectorNotMatchedMsgCount** | Pointer to **int64** | The number of guaranteed messages for which the Flow selector did not match. | [optional] 
**SessionName** | Pointer to **string** | The name of the Transacted Session for the Flow. | [optional] 
**StoreAndForwardAckedMsgCount** | Pointer to **int64** | The number of guaranteed messages that used store and forward delivery and are acknowledged by the consumer. | [optional] 
**TransportRetransmitMsgCount** | Pointer to **int64** | The number of guaranteed messages that were retransmitted at the transport layer as part of a single delivery attempt. Available since 2.18. | [optional] 
**UnackedMsgCount** | Pointer to **int64** | The number of guaranteed messages delivered but not yet acknowledged by the consumer. | [optional] 
**UsedWindowSize** | Pointer to **int64** | The number of guaranteed messages using the available window size. | [optional] 
**WindowClosedCount** | Pointer to **int64** | The number of times the window for guaranteed messages was filled and closed before an acknowledgement was received. | [optional] 
**WindowSize** | Pointer to **int64** | The number of outstanding guaranteed messages that can be transmitted over the Flow before an acknowledgement is received. | [optional] 

## Methods

### NewMsgVpnQueueTxFlow

`func NewMsgVpnQueueTxFlow() *MsgVpnQueueTxFlow`

NewMsgVpnQueueTxFlow instantiates a new MsgVpnQueueTxFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnQueueTxFlowWithDefaults

`func NewMsgVpnQueueTxFlowWithDefaults() *MsgVpnQueueTxFlow`

NewMsgVpnQueueTxFlowWithDefaults instantiates a new MsgVpnQueueTxFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAckedMsgCount

`func (o *MsgVpnQueueTxFlow) GetAckedMsgCount() int64`

GetAckedMsgCount returns the AckedMsgCount field if non-nil, zero value otherwise.

### GetAckedMsgCountOk

`func (o *MsgVpnQueueTxFlow) GetAckedMsgCountOk() (*int64, bool)`

GetAckedMsgCountOk returns a tuple with the AckedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckedMsgCount

`func (o *MsgVpnQueueTxFlow) SetAckedMsgCount(v int64)`

SetAckedMsgCount sets AckedMsgCount field to given value.

### HasAckedMsgCount

`func (o *MsgVpnQueueTxFlow) HasAckedMsgCount() bool`

HasAckedMsgCount returns a boolean if a field has been set.

### GetActivationTime

`func (o *MsgVpnQueueTxFlow) GetActivationTime() int32`

GetActivationTime returns the ActivationTime field if non-nil, zero value otherwise.

### GetActivationTimeOk

`func (o *MsgVpnQueueTxFlow) GetActivationTimeOk() (*int32, bool)`

GetActivationTimeOk returns a tuple with the ActivationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationTime

`func (o *MsgVpnQueueTxFlow) SetActivationTime(v int32)`

SetActivationTime sets ActivationTime field to given value.

### HasActivationTime

`func (o *MsgVpnQueueTxFlow) HasActivationTime() bool`

HasActivationTime returns a boolean if a field has been set.

### GetActivityState

`func (o *MsgVpnQueueTxFlow) GetActivityState() string`

GetActivityState returns the ActivityState field if non-nil, zero value otherwise.

### GetActivityStateOk

`func (o *MsgVpnQueueTxFlow) GetActivityStateOk() (*string, bool)`

GetActivityStateOk returns a tuple with the ActivityState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityState

`func (o *MsgVpnQueueTxFlow) SetActivityState(v string)`

SetActivityState sets ActivityState field to given value.

### HasActivityState

`func (o *MsgVpnQueueTxFlow) HasActivityState() bool`

HasActivityState returns a boolean if a field has been set.

### GetActivityUpdateState

`func (o *MsgVpnQueueTxFlow) GetActivityUpdateState() string`

GetActivityUpdateState returns the ActivityUpdateState field if non-nil, zero value otherwise.

### GetActivityUpdateStateOk

`func (o *MsgVpnQueueTxFlow) GetActivityUpdateStateOk() (*string, bool)`

GetActivityUpdateStateOk returns a tuple with the ActivityUpdateState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityUpdateState

`func (o *MsgVpnQueueTxFlow) SetActivityUpdateState(v string)`

SetActivityUpdateState sets ActivityUpdateState field to given value.

### HasActivityUpdateState

`func (o *MsgVpnQueueTxFlow) HasActivityUpdateState() bool`

HasActivityUpdateState returns a boolean if a field has been set.

### GetBindTime

`func (o *MsgVpnQueueTxFlow) GetBindTime() int32`

GetBindTime returns the BindTime field if non-nil, zero value otherwise.

### GetBindTimeOk

`func (o *MsgVpnQueueTxFlow) GetBindTimeOk() (*int32, bool)`

GetBindTimeOk returns a tuple with the BindTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindTime

`func (o *MsgVpnQueueTxFlow) SetBindTime(v int32)`

SetBindTime sets BindTime field to given value.

### HasBindTime

`func (o *MsgVpnQueueTxFlow) HasBindTime() bool`

HasBindTime returns a boolean if a field has been set.

### GetClientName

`func (o *MsgVpnQueueTxFlow) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *MsgVpnQueueTxFlow) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *MsgVpnQueueTxFlow) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *MsgVpnQueueTxFlow) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetConsumerRedeliveryRequestAllowed

`func (o *MsgVpnQueueTxFlow) GetConsumerRedeliveryRequestAllowed() bool`

GetConsumerRedeliveryRequestAllowed returns the ConsumerRedeliveryRequestAllowed field if non-nil, zero value otherwise.

### GetConsumerRedeliveryRequestAllowedOk

`func (o *MsgVpnQueueTxFlow) GetConsumerRedeliveryRequestAllowedOk() (*bool, bool)`

GetConsumerRedeliveryRequestAllowedOk returns a tuple with the ConsumerRedeliveryRequestAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerRedeliveryRequestAllowed

`func (o *MsgVpnQueueTxFlow) SetConsumerRedeliveryRequestAllowed(v bool)`

SetConsumerRedeliveryRequestAllowed sets ConsumerRedeliveryRequestAllowed field to given value.

### HasConsumerRedeliveryRequestAllowed

`func (o *MsgVpnQueueTxFlow) HasConsumerRedeliveryRequestAllowed() bool`

HasConsumerRedeliveryRequestAllowed returns a boolean if a field has been set.

### GetCutThroughAckedMsgCount

`func (o *MsgVpnQueueTxFlow) GetCutThroughAckedMsgCount() int64`

GetCutThroughAckedMsgCount returns the CutThroughAckedMsgCount field if non-nil, zero value otherwise.

### GetCutThroughAckedMsgCountOk

`func (o *MsgVpnQueueTxFlow) GetCutThroughAckedMsgCountOk() (*int64, bool)`

GetCutThroughAckedMsgCountOk returns a tuple with the CutThroughAckedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCutThroughAckedMsgCount

`func (o *MsgVpnQueueTxFlow) SetCutThroughAckedMsgCount(v int64)`

SetCutThroughAckedMsgCount sets CutThroughAckedMsgCount field to given value.

### HasCutThroughAckedMsgCount

`func (o *MsgVpnQueueTxFlow) HasCutThroughAckedMsgCount() bool`

HasCutThroughAckedMsgCount returns a boolean if a field has been set.

### GetDeliveryState

`func (o *MsgVpnQueueTxFlow) GetDeliveryState() string`

GetDeliveryState returns the DeliveryState field if non-nil, zero value otherwise.

### GetDeliveryStateOk

`func (o *MsgVpnQueueTxFlow) GetDeliveryStateOk() (*string, bool)`

GetDeliveryStateOk returns a tuple with the DeliveryState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryState

`func (o *MsgVpnQueueTxFlow) SetDeliveryState(v string)`

SetDeliveryState sets DeliveryState field to given value.

### HasDeliveryState

`func (o *MsgVpnQueueTxFlow) HasDeliveryState() bool`

HasDeliveryState returns a boolean if a field has been set.

### GetFlowId

`func (o *MsgVpnQueueTxFlow) GetFlowId() int64`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *MsgVpnQueueTxFlow) GetFlowIdOk() (*int64, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *MsgVpnQueueTxFlow) SetFlowId(v int64)`

SetFlowId sets FlowId field to given value.

### HasFlowId

`func (o *MsgVpnQueueTxFlow) HasFlowId() bool`

HasFlowId returns a boolean if a field has been set.

### GetHighestAckPendingMsgId

`func (o *MsgVpnQueueTxFlow) GetHighestAckPendingMsgId() int64`

GetHighestAckPendingMsgId returns the HighestAckPendingMsgId field if non-nil, zero value otherwise.

### GetHighestAckPendingMsgIdOk

`func (o *MsgVpnQueueTxFlow) GetHighestAckPendingMsgIdOk() (*int64, bool)`

GetHighestAckPendingMsgIdOk returns a tuple with the HighestAckPendingMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestAckPendingMsgId

`func (o *MsgVpnQueueTxFlow) SetHighestAckPendingMsgId(v int64)`

SetHighestAckPendingMsgId sets HighestAckPendingMsgId field to given value.

### HasHighestAckPendingMsgId

`func (o *MsgVpnQueueTxFlow) HasHighestAckPendingMsgId() bool`

HasHighestAckPendingMsgId returns a boolean if a field has been set.

### GetLastAckedMsgId

`func (o *MsgVpnQueueTxFlow) GetLastAckedMsgId() int64`

GetLastAckedMsgId returns the LastAckedMsgId field if non-nil, zero value otherwise.

### GetLastAckedMsgIdOk

`func (o *MsgVpnQueueTxFlow) GetLastAckedMsgIdOk() (*int64, bool)`

GetLastAckedMsgIdOk returns a tuple with the LastAckedMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAckedMsgId

`func (o *MsgVpnQueueTxFlow) SetLastAckedMsgId(v int64)`

SetLastAckedMsgId sets LastAckedMsgId field to given value.

### HasLastAckedMsgId

`func (o *MsgVpnQueueTxFlow) HasLastAckedMsgId() bool`

HasLastAckedMsgId returns a boolean if a field has been set.

### GetLastSelectorExaminedMsgId

`func (o *MsgVpnQueueTxFlow) GetLastSelectorExaminedMsgId() int64`

GetLastSelectorExaminedMsgId returns the LastSelectorExaminedMsgId field if non-nil, zero value otherwise.

### GetLastSelectorExaminedMsgIdOk

`func (o *MsgVpnQueueTxFlow) GetLastSelectorExaminedMsgIdOk() (*int64, bool)`

GetLastSelectorExaminedMsgIdOk returns a tuple with the LastSelectorExaminedMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSelectorExaminedMsgId

`func (o *MsgVpnQueueTxFlow) SetLastSelectorExaminedMsgId(v int64)`

SetLastSelectorExaminedMsgId sets LastSelectorExaminedMsgId field to given value.

### HasLastSelectorExaminedMsgId

`func (o *MsgVpnQueueTxFlow) HasLastSelectorExaminedMsgId() bool`

HasLastSelectorExaminedMsgId returns a boolean if a field has been set.

### GetLowestAckPendingMsgId

`func (o *MsgVpnQueueTxFlow) GetLowestAckPendingMsgId() int64`

GetLowestAckPendingMsgId returns the LowestAckPendingMsgId field if non-nil, zero value otherwise.

### GetLowestAckPendingMsgIdOk

`func (o *MsgVpnQueueTxFlow) GetLowestAckPendingMsgIdOk() (*int64, bool)`

GetLowestAckPendingMsgIdOk returns a tuple with the LowestAckPendingMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowestAckPendingMsgId

`func (o *MsgVpnQueueTxFlow) SetLowestAckPendingMsgId(v int64)`

SetLowestAckPendingMsgId sets LowestAckPendingMsgId field to given value.

### HasLowestAckPendingMsgId

`func (o *MsgVpnQueueTxFlow) HasLowestAckPendingMsgId() bool`

HasLowestAckPendingMsgId returns a boolean if a field has been set.

### GetMaxUnackedMsgsExceededMsgCount

`func (o *MsgVpnQueueTxFlow) GetMaxUnackedMsgsExceededMsgCount() int64`

GetMaxUnackedMsgsExceededMsgCount returns the MaxUnackedMsgsExceededMsgCount field if non-nil, zero value otherwise.

### GetMaxUnackedMsgsExceededMsgCountOk

`func (o *MsgVpnQueueTxFlow) GetMaxUnackedMsgsExceededMsgCountOk() (*int64, bool)`

GetMaxUnackedMsgsExceededMsgCountOk returns a tuple with the MaxUnackedMsgsExceededMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUnackedMsgsExceededMsgCount

`func (o *MsgVpnQueueTxFlow) SetMaxUnackedMsgsExceededMsgCount(v int64)`

SetMaxUnackedMsgsExceededMsgCount sets MaxUnackedMsgsExceededMsgCount field to given value.

### HasMaxUnackedMsgsExceededMsgCount

`func (o *MsgVpnQueueTxFlow) HasMaxUnackedMsgsExceededMsgCount() bool`

HasMaxUnackedMsgsExceededMsgCount returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnQueueTxFlow) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnQueueTxFlow) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnQueueTxFlow) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnQueueTxFlow) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetNoLocalDelivery

`func (o *MsgVpnQueueTxFlow) GetNoLocalDelivery() bool`

GetNoLocalDelivery returns the NoLocalDelivery field if non-nil, zero value otherwise.

### GetNoLocalDeliveryOk

`func (o *MsgVpnQueueTxFlow) GetNoLocalDeliveryOk() (*bool, bool)`

GetNoLocalDeliveryOk returns a tuple with the NoLocalDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoLocalDelivery

`func (o *MsgVpnQueueTxFlow) SetNoLocalDelivery(v bool)`

SetNoLocalDelivery sets NoLocalDelivery field to given value.

### HasNoLocalDelivery

`func (o *MsgVpnQueueTxFlow) HasNoLocalDelivery() bool`

HasNoLocalDelivery returns a boolean if a field has been set.

### GetQueueName

`func (o *MsgVpnQueueTxFlow) GetQueueName() string`

GetQueueName returns the QueueName field if non-nil, zero value otherwise.

### GetQueueNameOk

`func (o *MsgVpnQueueTxFlow) GetQueueNameOk() (*string, bool)`

GetQueueNameOk returns a tuple with the QueueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueName

`func (o *MsgVpnQueueTxFlow) SetQueueName(v string)`

SetQueueName sets QueueName field to given value.

### HasQueueName

`func (o *MsgVpnQueueTxFlow) HasQueueName() bool`

HasQueueName returns a boolean if a field has been set.

### GetRedeliveredMsgCount

`func (o *MsgVpnQueueTxFlow) GetRedeliveredMsgCount() int64`

GetRedeliveredMsgCount returns the RedeliveredMsgCount field if non-nil, zero value otherwise.

### GetRedeliveredMsgCountOk

`func (o *MsgVpnQueueTxFlow) GetRedeliveredMsgCountOk() (*int64, bool)`

GetRedeliveredMsgCountOk returns a tuple with the RedeliveredMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeliveredMsgCount

`func (o *MsgVpnQueueTxFlow) SetRedeliveredMsgCount(v int64)`

SetRedeliveredMsgCount sets RedeliveredMsgCount field to given value.

### HasRedeliveredMsgCount

`func (o *MsgVpnQueueTxFlow) HasRedeliveredMsgCount() bool`

HasRedeliveredMsgCount returns a boolean if a field has been set.

### GetRedeliveryRequestCount

`func (o *MsgVpnQueueTxFlow) GetRedeliveryRequestCount() int64`

GetRedeliveryRequestCount returns the RedeliveryRequestCount field if non-nil, zero value otherwise.

### GetRedeliveryRequestCountOk

`func (o *MsgVpnQueueTxFlow) GetRedeliveryRequestCountOk() (*int64, bool)`

GetRedeliveryRequestCountOk returns a tuple with the RedeliveryRequestCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeliveryRequestCount

`func (o *MsgVpnQueueTxFlow) SetRedeliveryRequestCount(v int64)`

SetRedeliveryRequestCount sets RedeliveryRequestCount field to given value.

### HasRedeliveryRequestCount

`func (o *MsgVpnQueueTxFlow) HasRedeliveryRequestCount() bool`

HasRedeliveryRequestCount returns a boolean if a field has been set.

### GetSelector

`func (o *MsgVpnQueueTxFlow) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *MsgVpnQueueTxFlow) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *MsgVpnQueueTxFlow) SetSelector(v string)`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *MsgVpnQueueTxFlow) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetSelectorExaminedMsgCount

`func (o *MsgVpnQueueTxFlow) GetSelectorExaminedMsgCount() int64`

GetSelectorExaminedMsgCount returns the SelectorExaminedMsgCount field if non-nil, zero value otherwise.

### GetSelectorExaminedMsgCountOk

`func (o *MsgVpnQueueTxFlow) GetSelectorExaminedMsgCountOk() (*int64, bool)`

GetSelectorExaminedMsgCountOk returns a tuple with the SelectorExaminedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectorExaminedMsgCount

`func (o *MsgVpnQueueTxFlow) SetSelectorExaminedMsgCount(v int64)`

SetSelectorExaminedMsgCount sets SelectorExaminedMsgCount field to given value.

### HasSelectorExaminedMsgCount

`func (o *MsgVpnQueueTxFlow) HasSelectorExaminedMsgCount() bool`

HasSelectorExaminedMsgCount returns a boolean if a field has been set.

### GetSelectorMatchedMsgCount

`func (o *MsgVpnQueueTxFlow) GetSelectorMatchedMsgCount() int64`

GetSelectorMatchedMsgCount returns the SelectorMatchedMsgCount field if non-nil, zero value otherwise.

### GetSelectorMatchedMsgCountOk

`func (o *MsgVpnQueueTxFlow) GetSelectorMatchedMsgCountOk() (*int64, bool)`

GetSelectorMatchedMsgCountOk returns a tuple with the SelectorMatchedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectorMatchedMsgCount

`func (o *MsgVpnQueueTxFlow) SetSelectorMatchedMsgCount(v int64)`

SetSelectorMatchedMsgCount sets SelectorMatchedMsgCount field to given value.

### HasSelectorMatchedMsgCount

`func (o *MsgVpnQueueTxFlow) HasSelectorMatchedMsgCount() bool`

HasSelectorMatchedMsgCount returns a boolean if a field has been set.

### GetSelectorNotMatchedMsgCount

`func (o *MsgVpnQueueTxFlow) GetSelectorNotMatchedMsgCount() int64`

GetSelectorNotMatchedMsgCount returns the SelectorNotMatchedMsgCount field if non-nil, zero value otherwise.

### GetSelectorNotMatchedMsgCountOk

`func (o *MsgVpnQueueTxFlow) GetSelectorNotMatchedMsgCountOk() (*int64, bool)`

GetSelectorNotMatchedMsgCountOk returns a tuple with the SelectorNotMatchedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectorNotMatchedMsgCount

`func (o *MsgVpnQueueTxFlow) SetSelectorNotMatchedMsgCount(v int64)`

SetSelectorNotMatchedMsgCount sets SelectorNotMatchedMsgCount field to given value.

### HasSelectorNotMatchedMsgCount

`func (o *MsgVpnQueueTxFlow) HasSelectorNotMatchedMsgCount() bool`

HasSelectorNotMatchedMsgCount returns a boolean if a field has been set.

### GetSessionName

`func (o *MsgVpnQueueTxFlow) GetSessionName() string`

GetSessionName returns the SessionName field if non-nil, zero value otherwise.

### GetSessionNameOk

`func (o *MsgVpnQueueTxFlow) GetSessionNameOk() (*string, bool)`

GetSessionNameOk returns a tuple with the SessionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionName

`func (o *MsgVpnQueueTxFlow) SetSessionName(v string)`

SetSessionName sets SessionName field to given value.

### HasSessionName

`func (o *MsgVpnQueueTxFlow) HasSessionName() bool`

HasSessionName returns a boolean if a field has been set.

### GetStoreAndForwardAckedMsgCount

`func (o *MsgVpnQueueTxFlow) GetStoreAndForwardAckedMsgCount() int64`

GetStoreAndForwardAckedMsgCount returns the StoreAndForwardAckedMsgCount field if non-nil, zero value otherwise.

### GetStoreAndForwardAckedMsgCountOk

`func (o *MsgVpnQueueTxFlow) GetStoreAndForwardAckedMsgCountOk() (*int64, bool)`

GetStoreAndForwardAckedMsgCountOk returns a tuple with the StoreAndForwardAckedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreAndForwardAckedMsgCount

`func (o *MsgVpnQueueTxFlow) SetStoreAndForwardAckedMsgCount(v int64)`

SetStoreAndForwardAckedMsgCount sets StoreAndForwardAckedMsgCount field to given value.

### HasStoreAndForwardAckedMsgCount

`func (o *MsgVpnQueueTxFlow) HasStoreAndForwardAckedMsgCount() bool`

HasStoreAndForwardAckedMsgCount returns a boolean if a field has been set.

### GetTransportRetransmitMsgCount

`func (o *MsgVpnQueueTxFlow) GetTransportRetransmitMsgCount() int64`

GetTransportRetransmitMsgCount returns the TransportRetransmitMsgCount field if non-nil, zero value otherwise.

### GetTransportRetransmitMsgCountOk

`func (o *MsgVpnQueueTxFlow) GetTransportRetransmitMsgCountOk() (*int64, bool)`

GetTransportRetransmitMsgCountOk returns a tuple with the TransportRetransmitMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportRetransmitMsgCount

`func (o *MsgVpnQueueTxFlow) SetTransportRetransmitMsgCount(v int64)`

SetTransportRetransmitMsgCount sets TransportRetransmitMsgCount field to given value.

### HasTransportRetransmitMsgCount

`func (o *MsgVpnQueueTxFlow) HasTransportRetransmitMsgCount() bool`

HasTransportRetransmitMsgCount returns a boolean if a field has been set.

### GetUnackedMsgCount

`func (o *MsgVpnQueueTxFlow) GetUnackedMsgCount() int64`

GetUnackedMsgCount returns the UnackedMsgCount field if non-nil, zero value otherwise.

### GetUnackedMsgCountOk

`func (o *MsgVpnQueueTxFlow) GetUnackedMsgCountOk() (*int64, bool)`

GetUnackedMsgCountOk returns a tuple with the UnackedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnackedMsgCount

`func (o *MsgVpnQueueTxFlow) SetUnackedMsgCount(v int64)`

SetUnackedMsgCount sets UnackedMsgCount field to given value.

### HasUnackedMsgCount

`func (o *MsgVpnQueueTxFlow) HasUnackedMsgCount() bool`

HasUnackedMsgCount returns a boolean if a field has been set.

### GetUsedWindowSize

`func (o *MsgVpnQueueTxFlow) GetUsedWindowSize() int64`

GetUsedWindowSize returns the UsedWindowSize field if non-nil, zero value otherwise.

### GetUsedWindowSizeOk

`func (o *MsgVpnQueueTxFlow) GetUsedWindowSizeOk() (*int64, bool)`

GetUsedWindowSizeOk returns a tuple with the UsedWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedWindowSize

`func (o *MsgVpnQueueTxFlow) SetUsedWindowSize(v int64)`

SetUsedWindowSize sets UsedWindowSize field to given value.

### HasUsedWindowSize

`func (o *MsgVpnQueueTxFlow) HasUsedWindowSize() bool`

HasUsedWindowSize returns a boolean if a field has been set.

### GetWindowClosedCount

`func (o *MsgVpnQueueTxFlow) GetWindowClosedCount() int64`

GetWindowClosedCount returns the WindowClosedCount field if non-nil, zero value otherwise.

### GetWindowClosedCountOk

`func (o *MsgVpnQueueTxFlow) GetWindowClosedCountOk() (*int64, bool)`

GetWindowClosedCountOk returns a tuple with the WindowClosedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowClosedCount

`func (o *MsgVpnQueueTxFlow) SetWindowClosedCount(v int64)`

SetWindowClosedCount sets WindowClosedCount field to given value.

### HasWindowClosedCount

`func (o *MsgVpnQueueTxFlow) HasWindowClosedCount() bool`

HasWindowClosedCount returns a boolean if a field has been set.

### GetWindowSize

`func (o *MsgVpnQueueTxFlow) GetWindowSize() int64`

GetWindowSize returns the WindowSize field if non-nil, zero value otherwise.

### GetWindowSizeOk

`func (o *MsgVpnQueueTxFlow) GetWindowSizeOk() (*int64, bool)`

GetWindowSizeOk returns a tuple with the WindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSize

`func (o *MsgVpnQueueTxFlow) SetWindowSize(v int64)`

SetWindowSize sets WindowSize field to given value.

### HasWindowSize

`func (o *MsgVpnQueueTxFlow) HasWindowSize() bool`

HasWindowSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


