# MsgVpnClientRxFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientName** | Pointer to **string** | The name of the Client. | [optional] 
**ConnectTime** | Pointer to **int32** | The timestamp of when the Flow from the Client connected. | [optional] 
**DestinationGroupErrorDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages from the Flow discarded due to a destination group error. | [optional] 
**DuplicateDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages from the Flow discarded due to being a duplicate. | [optional] 
**EndpointDisabledDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages from the Flow discarded due to an eligible endpoint destination being disabled. | [optional] 
**EndpointUsageExceededDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages from the Flow discarded due to an eligible endpoint destination having its maximum message spool usage exceeded. | [optional] 
**ErroredDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages from the Flow discarded due to errors being detected. | [optional] 
**FlowId** | Pointer to **int64** | The identifier (ID) of the flow. | [optional] 
**FlowName** | Pointer to **string** | The name of the Flow. | [optional] 
**GuaranteedMsgCount** | Pointer to **int64** | The number of guaranteed messages from the Flow. | [optional] 
**LastRxMsgId** | Pointer to **int64** | The identifier (ID) of the last message received on the Flow. | [optional] 
**LocalMsgCountExceededDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages from the Flow discarded due to the maximum number of messages allowed on the broker being exceeded. | [optional] 
**LowPriorityMsgCongestionDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages from the Flow discarded due to congestion of low priority messages. | [optional] 
**MaxMsgSizeExceededDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages from the Flow discarded due to the maximum allowed message size being exceeded. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**NoEligibleDestinationsDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages from the Flow discarded due to there being no eligible endpoint destination. | [optional] 
**NoLocalDeliveryDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages from the Flow discarded due to no local delivery being requested. | [optional] 
**NotCompatibleWithForwardingModeDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages from the Flow discarded due to being incompatible with the forwarding mode of an eligible endpoint destination. | [optional] 
**OutOfOrderDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages from the Flow discarded due to being received out of order. | [optional] 
**PublishAclDeniedDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages from the Flow discarded due to being denied by the access control list (ACL) profile for the published topic. | [optional] 
**PublisherId** | Pointer to **int64** | The identifier (ID) of the publisher for the Flow. | [optional] 
**QueueNotFoundDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages from the Flow discarded due to the destination queue not being found. | [optional] 
**ReplicationStandbyDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages from the Flow discarded due to the Message VPN being in the replication standby state. | [optional] 
**SessionName** | Pointer to **string** | The name of the transacted session on the Flow. | [optional] 
**SmfTtlExceededDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages from the Flow discarded due to the message time-to-live (TTL) count being exceeded. The message TTL count is the maximum number of times the message can cross a bridge between Message VPNs. | [optional] 
**SpoolFileLimitExceededDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages from the Flow discarded due to all available message spool file resources being used. | [optional] 
**SpoolNotReadyDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages from the Flow discarded due to the message spool being not ready. | [optional] 
**SpoolToAdbFailDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages from the Flow discarded due to a failure while spooling to the Assured Delivery Blade (ADB). | [optional] 
**SpoolToDiskFailDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages from the Flow discarded due to a failure while spooling to the disk. | [optional] 
**SpoolUsageExceededDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages from the Flow discarded due to the maximum message spool usage being exceeded. | [optional] 
**SyncReplicationIneligibleDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages from the Flow discarded due to synchronous replication being ineligible. | [optional] 
**UserProfileDeniedGuaranteedDiscardedMsgCount** | Pointer to **int64** | The number of guaranteed messages from the Flow discarded due to being denied by the client profile. | [optional] 
**WindowSize** | Pointer to **int32** | The size of the window used for guaranteed messages sent on the Flow, in messages. | [optional] 

## Methods

### NewMsgVpnClientRxFlow

`func NewMsgVpnClientRxFlow() *MsgVpnClientRxFlow`

NewMsgVpnClientRxFlow instantiates a new MsgVpnClientRxFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnClientRxFlowWithDefaults

`func NewMsgVpnClientRxFlowWithDefaults() *MsgVpnClientRxFlow`

NewMsgVpnClientRxFlowWithDefaults instantiates a new MsgVpnClientRxFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientName

`func (o *MsgVpnClientRxFlow) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *MsgVpnClientRxFlow) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *MsgVpnClientRxFlow) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *MsgVpnClientRxFlow) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetConnectTime

`func (o *MsgVpnClientRxFlow) GetConnectTime() int32`

GetConnectTime returns the ConnectTime field if non-nil, zero value otherwise.

### GetConnectTimeOk

`func (o *MsgVpnClientRxFlow) GetConnectTimeOk() (*int32, bool)`

GetConnectTimeOk returns a tuple with the ConnectTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectTime

`func (o *MsgVpnClientRxFlow) SetConnectTime(v int32)`

SetConnectTime sets ConnectTime field to given value.

### HasConnectTime

`func (o *MsgVpnClientRxFlow) HasConnectTime() bool`

HasConnectTime returns a boolean if a field has been set.

### GetDestinationGroupErrorDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) GetDestinationGroupErrorDiscardedMsgCount() int64`

GetDestinationGroupErrorDiscardedMsgCount returns the DestinationGroupErrorDiscardedMsgCount field if non-nil, zero value otherwise.

### GetDestinationGroupErrorDiscardedMsgCountOk

`func (o *MsgVpnClientRxFlow) GetDestinationGroupErrorDiscardedMsgCountOk() (*int64, bool)`

GetDestinationGroupErrorDiscardedMsgCountOk returns a tuple with the DestinationGroupErrorDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationGroupErrorDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) SetDestinationGroupErrorDiscardedMsgCount(v int64)`

SetDestinationGroupErrorDiscardedMsgCount sets DestinationGroupErrorDiscardedMsgCount field to given value.

### HasDestinationGroupErrorDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) HasDestinationGroupErrorDiscardedMsgCount() bool`

HasDestinationGroupErrorDiscardedMsgCount returns a boolean if a field has been set.

### GetDuplicateDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) GetDuplicateDiscardedMsgCount() int64`

GetDuplicateDiscardedMsgCount returns the DuplicateDiscardedMsgCount field if non-nil, zero value otherwise.

### GetDuplicateDiscardedMsgCountOk

`func (o *MsgVpnClientRxFlow) GetDuplicateDiscardedMsgCountOk() (*int64, bool)`

GetDuplicateDiscardedMsgCountOk returns a tuple with the DuplicateDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicateDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) SetDuplicateDiscardedMsgCount(v int64)`

SetDuplicateDiscardedMsgCount sets DuplicateDiscardedMsgCount field to given value.

### HasDuplicateDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) HasDuplicateDiscardedMsgCount() bool`

HasDuplicateDiscardedMsgCount returns a boolean if a field has been set.

### GetEndpointDisabledDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) GetEndpointDisabledDiscardedMsgCount() int64`

GetEndpointDisabledDiscardedMsgCount returns the EndpointDisabledDiscardedMsgCount field if non-nil, zero value otherwise.

### GetEndpointDisabledDiscardedMsgCountOk

`func (o *MsgVpnClientRxFlow) GetEndpointDisabledDiscardedMsgCountOk() (*int64, bool)`

GetEndpointDisabledDiscardedMsgCountOk returns a tuple with the EndpointDisabledDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointDisabledDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) SetEndpointDisabledDiscardedMsgCount(v int64)`

SetEndpointDisabledDiscardedMsgCount sets EndpointDisabledDiscardedMsgCount field to given value.

### HasEndpointDisabledDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) HasEndpointDisabledDiscardedMsgCount() bool`

HasEndpointDisabledDiscardedMsgCount returns a boolean if a field has been set.

### GetEndpointUsageExceededDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) GetEndpointUsageExceededDiscardedMsgCount() int64`

GetEndpointUsageExceededDiscardedMsgCount returns the EndpointUsageExceededDiscardedMsgCount field if non-nil, zero value otherwise.

### GetEndpointUsageExceededDiscardedMsgCountOk

`func (o *MsgVpnClientRxFlow) GetEndpointUsageExceededDiscardedMsgCountOk() (*int64, bool)`

GetEndpointUsageExceededDiscardedMsgCountOk returns a tuple with the EndpointUsageExceededDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointUsageExceededDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) SetEndpointUsageExceededDiscardedMsgCount(v int64)`

SetEndpointUsageExceededDiscardedMsgCount sets EndpointUsageExceededDiscardedMsgCount field to given value.

### HasEndpointUsageExceededDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) HasEndpointUsageExceededDiscardedMsgCount() bool`

HasEndpointUsageExceededDiscardedMsgCount returns a boolean if a field has been set.

### GetErroredDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) GetErroredDiscardedMsgCount() int64`

GetErroredDiscardedMsgCount returns the ErroredDiscardedMsgCount field if non-nil, zero value otherwise.

### GetErroredDiscardedMsgCountOk

`func (o *MsgVpnClientRxFlow) GetErroredDiscardedMsgCountOk() (*int64, bool)`

GetErroredDiscardedMsgCountOk returns a tuple with the ErroredDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErroredDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) SetErroredDiscardedMsgCount(v int64)`

SetErroredDiscardedMsgCount sets ErroredDiscardedMsgCount field to given value.

### HasErroredDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) HasErroredDiscardedMsgCount() bool`

HasErroredDiscardedMsgCount returns a boolean if a field has been set.

### GetFlowId

`func (o *MsgVpnClientRxFlow) GetFlowId() int64`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *MsgVpnClientRxFlow) GetFlowIdOk() (*int64, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *MsgVpnClientRxFlow) SetFlowId(v int64)`

SetFlowId sets FlowId field to given value.

### HasFlowId

`func (o *MsgVpnClientRxFlow) HasFlowId() bool`

HasFlowId returns a boolean if a field has been set.

### GetFlowName

`func (o *MsgVpnClientRxFlow) GetFlowName() string`

GetFlowName returns the FlowName field if non-nil, zero value otherwise.

### GetFlowNameOk

`func (o *MsgVpnClientRxFlow) GetFlowNameOk() (*string, bool)`

GetFlowNameOk returns a tuple with the FlowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowName

`func (o *MsgVpnClientRxFlow) SetFlowName(v string)`

SetFlowName sets FlowName field to given value.

### HasFlowName

`func (o *MsgVpnClientRxFlow) HasFlowName() bool`

HasFlowName returns a boolean if a field has been set.

### GetGuaranteedMsgCount

`func (o *MsgVpnClientRxFlow) GetGuaranteedMsgCount() int64`

GetGuaranteedMsgCount returns the GuaranteedMsgCount field if non-nil, zero value otherwise.

### GetGuaranteedMsgCountOk

`func (o *MsgVpnClientRxFlow) GetGuaranteedMsgCountOk() (*int64, bool)`

GetGuaranteedMsgCountOk returns a tuple with the GuaranteedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuaranteedMsgCount

`func (o *MsgVpnClientRxFlow) SetGuaranteedMsgCount(v int64)`

SetGuaranteedMsgCount sets GuaranteedMsgCount field to given value.

### HasGuaranteedMsgCount

`func (o *MsgVpnClientRxFlow) HasGuaranteedMsgCount() bool`

HasGuaranteedMsgCount returns a boolean if a field has been set.

### GetLastRxMsgId

`func (o *MsgVpnClientRxFlow) GetLastRxMsgId() int64`

GetLastRxMsgId returns the LastRxMsgId field if non-nil, zero value otherwise.

### GetLastRxMsgIdOk

`func (o *MsgVpnClientRxFlow) GetLastRxMsgIdOk() (*int64, bool)`

GetLastRxMsgIdOk returns a tuple with the LastRxMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRxMsgId

`func (o *MsgVpnClientRxFlow) SetLastRxMsgId(v int64)`

SetLastRxMsgId sets LastRxMsgId field to given value.

### HasLastRxMsgId

`func (o *MsgVpnClientRxFlow) HasLastRxMsgId() bool`

HasLastRxMsgId returns a boolean if a field has been set.

### GetLocalMsgCountExceededDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) GetLocalMsgCountExceededDiscardedMsgCount() int64`

GetLocalMsgCountExceededDiscardedMsgCount returns the LocalMsgCountExceededDiscardedMsgCount field if non-nil, zero value otherwise.

### GetLocalMsgCountExceededDiscardedMsgCountOk

`func (o *MsgVpnClientRxFlow) GetLocalMsgCountExceededDiscardedMsgCountOk() (*int64, bool)`

GetLocalMsgCountExceededDiscardedMsgCountOk returns a tuple with the LocalMsgCountExceededDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalMsgCountExceededDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) SetLocalMsgCountExceededDiscardedMsgCount(v int64)`

SetLocalMsgCountExceededDiscardedMsgCount sets LocalMsgCountExceededDiscardedMsgCount field to given value.

### HasLocalMsgCountExceededDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) HasLocalMsgCountExceededDiscardedMsgCount() bool`

HasLocalMsgCountExceededDiscardedMsgCount returns a boolean if a field has been set.

### GetLowPriorityMsgCongestionDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) GetLowPriorityMsgCongestionDiscardedMsgCount() int64`

GetLowPriorityMsgCongestionDiscardedMsgCount returns the LowPriorityMsgCongestionDiscardedMsgCount field if non-nil, zero value otherwise.

### GetLowPriorityMsgCongestionDiscardedMsgCountOk

`func (o *MsgVpnClientRxFlow) GetLowPriorityMsgCongestionDiscardedMsgCountOk() (*int64, bool)`

GetLowPriorityMsgCongestionDiscardedMsgCountOk returns a tuple with the LowPriorityMsgCongestionDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowPriorityMsgCongestionDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) SetLowPriorityMsgCongestionDiscardedMsgCount(v int64)`

SetLowPriorityMsgCongestionDiscardedMsgCount sets LowPriorityMsgCongestionDiscardedMsgCount field to given value.

### HasLowPriorityMsgCongestionDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) HasLowPriorityMsgCongestionDiscardedMsgCount() bool`

HasLowPriorityMsgCongestionDiscardedMsgCount returns a boolean if a field has been set.

### GetMaxMsgSizeExceededDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) GetMaxMsgSizeExceededDiscardedMsgCount() int64`

GetMaxMsgSizeExceededDiscardedMsgCount returns the MaxMsgSizeExceededDiscardedMsgCount field if non-nil, zero value otherwise.

### GetMaxMsgSizeExceededDiscardedMsgCountOk

`func (o *MsgVpnClientRxFlow) GetMaxMsgSizeExceededDiscardedMsgCountOk() (*int64, bool)`

GetMaxMsgSizeExceededDiscardedMsgCountOk returns a tuple with the MaxMsgSizeExceededDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMsgSizeExceededDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) SetMaxMsgSizeExceededDiscardedMsgCount(v int64)`

SetMaxMsgSizeExceededDiscardedMsgCount sets MaxMsgSizeExceededDiscardedMsgCount field to given value.

### HasMaxMsgSizeExceededDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) HasMaxMsgSizeExceededDiscardedMsgCount() bool`

HasMaxMsgSizeExceededDiscardedMsgCount returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnClientRxFlow) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnClientRxFlow) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnClientRxFlow) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnClientRxFlow) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetNoEligibleDestinationsDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) GetNoEligibleDestinationsDiscardedMsgCount() int64`

GetNoEligibleDestinationsDiscardedMsgCount returns the NoEligibleDestinationsDiscardedMsgCount field if non-nil, zero value otherwise.

### GetNoEligibleDestinationsDiscardedMsgCountOk

`func (o *MsgVpnClientRxFlow) GetNoEligibleDestinationsDiscardedMsgCountOk() (*int64, bool)`

GetNoEligibleDestinationsDiscardedMsgCountOk returns a tuple with the NoEligibleDestinationsDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoEligibleDestinationsDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) SetNoEligibleDestinationsDiscardedMsgCount(v int64)`

SetNoEligibleDestinationsDiscardedMsgCount sets NoEligibleDestinationsDiscardedMsgCount field to given value.

### HasNoEligibleDestinationsDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) HasNoEligibleDestinationsDiscardedMsgCount() bool`

HasNoEligibleDestinationsDiscardedMsgCount returns a boolean if a field has been set.

### GetNoLocalDeliveryDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) GetNoLocalDeliveryDiscardedMsgCount() int64`

GetNoLocalDeliveryDiscardedMsgCount returns the NoLocalDeliveryDiscardedMsgCount field if non-nil, zero value otherwise.

### GetNoLocalDeliveryDiscardedMsgCountOk

`func (o *MsgVpnClientRxFlow) GetNoLocalDeliveryDiscardedMsgCountOk() (*int64, bool)`

GetNoLocalDeliveryDiscardedMsgCountOk returns a tuple with the NoLocalDeliveryDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoLocalDeliveryDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) SetNoLocalDeliveryDiscardedMsgCount(v int64)`

SetNoLocalDeliveryDiscardedMsgCount sets NoLocalDeliveryDiscardedMsgCount field to given value.

### HasNoLocalDeliveryDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) HasNoLocalDeliveryDiscardedMsgCount() bool`

HasNoLocalDeliveryDiscardedMsgCount returns a boolean if a field has been set.

### GetNotCompatibleWithForwardingModeDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) GetNotCompatibleWithForwardingModeDiscardedMsgCount() int64`

GetNotCompatibleWithForwardingModeDiscardedMsgCount returns the NotCompatibleWithForwardingModeDiscardedMsgCount field if non-nil, zero value otherwise.

### GetNotCompatibleWithForwardingModeDiscardedMsgCountOk

`func (o *MsgVpnClientRxFlow) GetNotCompatibleWithForwardingModeDiscardedMsgCountOk() (*int64, bool)`

GetNotCompatibleWithForwardingModeDiscardedMsgCountOk returns a tuple with the NotCompatibleWithForwardingModeDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotCompatibleWithForwardingModeDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) SetNotCompatibleWithForwardingModeDiscardedMsgCount(v int64)`

SetNotCompatibleWithForwardingModeDiscardedMsgCount sets NotCompatibleWithForwardingModeDiscardedMsgCount field to given value.

### HasNotCompatibleWithForwardingModeDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) HasNotCompatibleWithForwardingModeDiscardedMsgCount() bool`

HasNotCompatibleWithForwardingModeDiscardedMsgCount returns a boolean if a field has been set.

### GetOutOfOrderDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) GetOutOfOrderDiscardedMsgCount() int64`

GetOutOfOrderDiscardedMsgCount returns the OutOfOrderDiscardedMsgCount field if non-nil, zero value otherwise.

### GetOutOfOrderDiscardedMsgCountOk

`func (o *MsgVpnClientRxFlow) GetOutOfOrderDiscardedMsgCountOk() (*int64, bool)`

GetOutOfOrderDiscardedMsgCountOk returns a tuple with the OutOfOrderDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutOfOrderDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) SetOutOfOrderDiscardedMsgCount(v int64)`

SetOutOfOrderDiscardedMsgCount sets OutOfOrderDiscardedMsgCount field to given value.

### HasOutOfOrderDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) HasOutOfOrderDiscardedMsgCount() bool`

HasOutOfOrderDiscardedMsgCount returns a boolean if a field has been set.

### GetPublishAclDeniedDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) GetPublishAclDeniedDiscardedMsgCount() int64`

GetPublishAclDeniedDiscardedMsgCount returns the PublishAclDeniedDiscardedMsgCount field if non-nil, zero value otherwise.

### GetPublishAclDeniedDiscardedMsgCountOk

`func (o *MsgVpnClientRxFlow) GetPublishAclDeniedDiscardedMsgCountOk() (*int64, bool)`

GetPublishAclDeniedDiscardedMsgCountOk returns a tuple with the PublishAclDeniedDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishAclDeniedDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) SetPublishAclDeniedDiscardedMsgCount(v int64)`

SetPublishAclDeniedDiscardedMsgCount sets PublishAclDeniedDiscardedMsgCount field to given value.

### HasPublishAclDeniedDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) HasPublishAclDeniedDiscardedMsgCount() bool`

HasPublishAclDeniedDiscardedMsgCount returns a boolean if a field has been set.

### GetPublisherId

`func (o *MsgVpnClientRxFlow) GetPublisherId() int64`

GetPublisherId returns the PublisherId field if non-nil, zero value otherwise.

### GetPublisherIdOk

`func (o *MsgVpnClientRxFlow) GetPublisherIdOk() (*int64, bool)`

GetPublisherIdOk returns a tuple with the PublisherId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherId

`func (o *MsgVpnClientRxFlow) SetPublisherId(v int64)`

SetPublisherId sets PublisherId field to given value.

### HasPublisherId

`func (o *MsgVpnClientRxFlow) HasPublisherId() bool`

HasPublisherId returns a boolean if a field has been set.

### GetQueueNotFoundDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) GetQueueNotFoundDiscardedMsgCount() int64`

GetQueueNotFoundDiscardedMsgCount returns the QueueNotFoundDiscardedMsgCount field if non-nil, zero value otherwise.

### GetQueueNotFoundDiscardedMsgCountOk

`func (o *MsgVpnClientRxFlow) GetQueueNotFoundDiscardedMsgCountOk() (*int64, bool)`

GetQueueNotFoundDiscardedMsgCountOk returns a tuple with the QueueNotFoundDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueNotFoundDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) SetQueueNotFoundDiscardedMsgCount(v int64)`

SetQueueNotFoundDiscardedMsgCount sets QueueNotFoundDiscardedMsgCount field to given value.

### HasQueueNotFoundDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) HasQueueNotFoundDiscardedMsgCount() bool`

HasQueueNotFoundDiscardedMsgCount returns a boolean if a field has been set.

### GetReplicationStandbyDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) GetReplicationStandbyDiscardedMsgCount() int64`

GetReplicationStandbyDiscardedMsgCount returns the ReplicationStandbyDiscardedMsgCount field if non-nil, zero value otherwise.

### GetReplicationStandbyDiscardedMsgCountOk

`func (o *MsgVpnClientRxFlow) GetReplicationStandbyDiscardedMsgCountOk() (*int64, bool)`

GetReplicationStandbyDiscardedMsgCountOk returns a tuple with the ReplicationStandbyDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationStandbyDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) SetReplicationStandbyDiscardedMsgCount(v int64)`

SetReplicationStandbyDiscardedMsgCount sets ReplicationStandbyDiscardedMsgCount field to given value.

### HasReplicationStandbyDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) HasReplicationStandbyDiscardedMsgCount() bool`

HasReplicationStandbyDiscardedMsgCount returns a boolean if a field has been set.

### GetSessionName

`func (o *MsgVpnClientRxFlow) GetSessionName() string`

GetSessionName returns the SessionName field if non-nil, zero value otherwise.

### GetSessionNameOk

`func (o *MsgVpnClientRxFlow) GetSessionNameOk() (*string, bool)`

GetSessionNameOk returns a tuple with the SessionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionName

`func (o *MsgVpnClientRxFlow) SetSessionName(v string)`

SetSessionName sets SessionName field to given value.

### HasSessionName

`func (o *MsgVpnClientRxFlow) HasSessionName() bool`

HasSessionName returns a boolean if a field has been set.

### GetSmfTtlExceededDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) GetSmfTtlExceededDiscardedMsgCount() int64`

GetSmfTtlExceededDiscardedMsgCount returns the SmfTtlExceededDiscardedMsgCount field if non-nil, zero value otherwise.

### GetSmfTtlExceededDiscardedMsgCountOk

`func (o *MsgVpnClientRxFlow) GetSmfTtlExceededDiscardedMsgCountOk() (*int64, bool)`

GetSmfTtlExceededDiscardedMsgCountOk returns a tuple with the SmfTtlExceededDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmfTtlExceededDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) SetSmfTtlExceededDiscardedMsgCount(v int64)`

SetSmfTtlExceededDiscardedMsgCount sets SmfTtlExceededDiscardedMsgCount field to given value.

### HasSmfTtlExceededDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) HasSmfTtlExceededDiscardedMsgCount() bool`

HasSmfTtlExceededDiscardedMsgCount returns a boolean if a field has been set.

### GetSpoolFileLimitExceededDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) GetSpoolFileLimitExceededDiscardedMsgCount() int64`

GetSpoolFileLimitExceededDiscardedMsgCount returns the SpoolFileLimitExceededDiscardedMsgCount field if non-nil, zero value otherwise.

### GetSpoolFileLimitExceededDiscardedMsgCountOk

`func (o *MsgVpnClientRxFlow) GetSpoolFileLimitExceededDiscardedMsgCountOk() (*int64, bool)`

GetSpoolFileLimitExceededDiscardedMsgCountOk returns a tuple with the SpoolFileLimitExceededDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpoolFileLimitExceededDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) SetSpoolFileLimitExceededDiscardedMsgCount(v int64)`

SetSpoolFileLimitExceededDiscardedMsgCount sets SpoolFileLimitExceededDiscardedMsgCount field to given value.

### HasSpoolFileLimitExceededDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) HasSpoolFileLimitExceededDiscardedMsgCount() bool`

HasSpoolFileLimitExceededDiscardedMsgCount returns a boolean if a field has been set.

### GetSpoolNotReadyDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) GetSpoolNotReadyDiscardedMsgCount() int64`

GetSpoolNotReadyDiscardedMsgCount returns the SpoolNotReadyDiscardedMsgCount field if non-nil, zero value otherwise.

### GetSpoolNotReadyDiscardedMsgCountOk

`func (o *MsgVpnClientRxFlow) GetSpoolNotReadyDiscardedMsgCountOk() (*int64, bool)`

GetSpoolNotReadyDiscardedMsgCountOk returns a tuple with the SpoolNotReadyDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpoolNotReadyDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) SetSpoolNotReadyDiscardedMsgCount(v int64)`

SetSpoolNotReadyDiscardedMsgCount sets SpoolNotReadyDiscardedMsgCount field to given value.

### HasSpoolNotReadyDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) HasSpoolNotReadyDiscardedMsgCount() bool`

HasSpoolNotReadyDiscardedMsgCount returns a boolean if a field has been set.

### GetSpoolToAdbFailDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) GetSpoolToAdbFailDiscardedMsgCount() int64`

GetSpoolToAdbFailDiscardedMsgCount returns the SpoolToAdbFailDiscardedMsgCount field if non-nil, zero value otherwise.

### GetSpoolToAdbFailDiscardedMsgCountOk

`func (o *MsgVpnClientRxFlow) GetSpoolToAdbFailDiscardedMsgCountOk() (*int64, bool)`

GetSpoolToAdbFailDiscardedMsgCountOk returns a tuple with the SpoolToAdbFailDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpoolToAdbFailDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) SetSpoolToAdbFailDiscardedMsgCount(v int64)`

SetSpoolToAdbFailDiscardedMsgCount sets SpoolToAdbFailDiscardedMsgCount field to given value.

### HasSpoolToAdbFailDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) HasSpoolToAdbFailDiscardedMsgCount() bool`

HasSpoolToAdbFailDiscardedMsgCount returns a boolean if a field has been set.

### GetSpoolToDiskFailDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) GetSpoolToDiskFailDiscardedMsgCount() int64`

GetSpoolToDiskFailDiscardedMsgCount returns the SpoolToDiskFailDiscardedMsgCount field if non-nil, zero value otherwise.

### GetSpoolToDiskFailDiscardedMsgCountOk

`func (o *MsgVpnClientRxFlow) GetSpoolToDiskFailDiscardedMsgCountOk() (*int64, bool)`

GetSpoolToDiskFailDiscardedMsgCountOk returns a tuple with the SpoolToDiskFailDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpoolToDiskFailDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) SetSpoolToDiskFailDiscardedMsgCount(v int64)`

SetSpoolToDiskFailDiscardedMsgCount sets SpoolToDiskFailDiscardedMsgCount field to given value.

### HasSpoolToDiskFailDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) HasSpoolToDiskFailDiscardedMsgCount() bool`

HasSpoolToDiskFailDiscardedMsgCount returns a boolean if a field has been set.

### GetSpoolUsageExceededDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) GetSpoolUsageExceededDiscardedMsgCount() int64`

GetSpoolUsageExceededDiscardedMsgCount returns the SpoolUsageExceededDiscardedMsgCount field if non-nil, zero value otherwise.

### GetSpoolUsageExceededDiscardedMsgCountOk

`func (o *MsgVpnClientRxFlow) GetSpoolUsageExceededDiscardedMsgCountOk() (*int64, bool)`

GetSpoolUsageExceededDiscardedMsgCountOk returns a tuple with the SpoolUsageExceededDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpoolUsageExceededDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) SetSpoolUsageExceededDiscardedMsgCount(v int64)`

SetSpoolUsageExceededDiscardedMsgCount sets SpoolUsageExceededDiscardedMsgCount field to given value.

### HasSpoolUsageExceededDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) HasSpoolUsageExceededDiscardedMsgCount() bool`

HasSpoolUsageExceededDiscardedMsgCount returns a boolean if a field has been set.

### GetSyncReplicationIneligibleDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) GetSyncReplicationIneligibleDiscardedMsgCount() int64`

GetSyncReplicationIneligibleDiscardedMsgCount returns the SyncReplicationIneligibleDiscardedMsgCount field if non-nil, zero value otherwise.

### GetSyncReplicationIneligibleDiscardedMsgCountOk

`func (o *MsgVpnClientRxFlow) GetSyncReplicationIneligibleDiscardedMsgCountOk() (*int64, bool)`

GetSyncReplicationIneligibleDiscardedMsgCountOk returns a tuple with the SyncReplicationIneligibleDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncReplicationIneligibleDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) SetSyncReplicationIneligibleDiscardedMsgCount(v int64)`

SetSyncReplicationIneligibleDiscardedMsgCount sets SyncReplicationIneligibleDiscardedMsgCount field to given value.

### HasSyncReplicationIneligibleDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) HasSyncReplicationIneligibleDiscardedMsgCount() bool`

HasSyncReplicationIneligibleDiscardedMsgCount returns a boolean if a field has been set.

### GetUserProfileDeniedGuaranteedDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) GetUserProfileDeniedGuaranteedDiscardedMsgCount() int64`

GetUserProfileDeniedGuaranteedDiscardedMsgCount returns the UserProfileDeniedGuaranteedDiscardedMsgCount field if non-nil, zero value otherwise.

### GetUserProfileDeniedGuaranteedDiscardedMsgCountOk

`func (o *MsgVpnClientRxFlow) GetUserProfileDeniedGuaranteedDiscardedMsgCountOk() (*int64, bool)`

GetUserProfileDeniedGuaranteedDiscardedMsgCountOk returns a tuple with the UserProfileDeniedGuaranteedDiscardedMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProfileDeniedGuaranteedDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) SetUserProfileDeniedGuaranteedDiscardedMsgCount(v int64)`

SetUserProfileDeniedGuaranteedDiscardedMsgCount sets UserProfileDeniedGuaranteedDiscardedMsgCount field to given value.

### HasUserProfileDeniedGuaranteedDiscardedMsgCount

`func (o *MsgVpnClientRxFlow) HasUserProfileDeniedGuaranteedDiscardedMsgCount() bool`

HasUserProfileDeniedGuaranteedDiscardedMsgCount returns a boolean if a field has been set.

### GetWindowSize

`func (o *MsgVpnClientRxFlow) GetWindowSize() int32`

GetWindowSize returns the WindowSize field if non-nil, zero value otherwise.

### GetWindowSizeOk

`func (o *MsgVpnClientRxFlow) GetWindowSizeOk() (*int32, bool)`

GetWindowSizeOk returns a tuple with the WindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowSize

`func (o *MsgVpnClientRxFlow) SetWindowSize(v int32)`

SetWindowSize sets WindowSize field to given value.

### HasWindowSize

`func (o *MsgVpnClientRxFlow) HasWindowSize() bool`

HasWindowSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


