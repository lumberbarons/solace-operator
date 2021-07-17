# MsgVpnRestDeliveryPointQueueBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GatewayReplaceTargetAuthorityEnabled** | Pointer to **bool** | Indicates whether the authority for the request-target is replaced with that configured for the REST Consumer remote. | [optional] 
**LastFailureReason** | Pointer to **string** | The reason for the last REST Delivery Point queue binding failure. | [optional] 
**LastFailureTime** | Pointer to **int32** | The timestamp of the last REST Delivery Point queue binding failure. This value represents the number of seconds since 1970-01-01 00:00:00 UTC (Unix time). | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**PostRequestTarget** | Pointer to **string** | The request-target string being used when sending requests to a REST Consumer. | [optional] 
**QueueBindingName** | Pointer to **string** | The name of a queue in the Message VPN. | [optional] 
**RestDeliveryPointName** | Pointer to **string** | The name of the REST Delivery Point. | [optional] 
**Up** | Pointer to **bool** | Indicates whether the operational state of the REST Delivery Point queue binding is up. | [optional] 
**Uptime** | Pointer to **int64** | The amount of time in seconds since the REST Delivery Point queue binding was up. | [optional] 

## Methods

### NewMsgVpnRestDeliveryPointQueueBinding

`func NewMsgVpnRestDeliveryPointQueueBinding() *MsgVpnRestDeliveryPointQueueBinding`

NewMsgVpnRestDeliveryPointQueueBinding instantiates a new MsgVpnRestDeliveryPointQueueBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnRestDeliveryPointQueueBindingWithDefaults

`func NewMsgVpnRestDeliveryPointQueueBindingWithDefaults() *MsgVpnRestDeliveryPointQueueBinding`

NewMsgVpnRestDeliveryPointQueueBindingWithDefaults instantiates a new MsgVpnRestDeliveryPointQueueBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGatewayReplaceTargetAuthorityEnabled

`func (o *MsgVpnRestDeliveryPointQueueBinding) GetGatewayReplaceTargetAuthorityEnabled() bool`

GetGatewayReplaceTargetAuthorityEnabled returns the GatewayReplaceTargetAuthorityEnabled field if non-nil, zero value otherwise.

### GetGatewayReplaceTargetAuthorityEnabledOk

`func (o *MsgVpnRestDeliveryPointQueueBinding) GetGatewayReplaceTargetAuthorityEnabledOk() (*bool, bool)`

GetGatewayReplaceTargetAuthorityEnabledOk returns a tuple with the GatewayReplaceTargetAuthorityEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayReplaceTargetAuthorityEnabled

`func (o *MsgVpnRestDeliveryPointQueueBinding) SetGatewayReplaceTargetAuthorityEnabled(v bool)`

SetGatewayReplaceTargetAuthorityEnabled sets GatewayReplaceTargetAuthorityEnabled field to given value.

### HasGatewayReplaceTargetAuthorityEnabled

`func (o *MsgVpnRestDeliveryPointQueueBinding) HasGatewayReplaceTargetAuthorityEnabled() bool`

HasGatewayReplaceTargetAuthorityEnabled returns a boolean if a field has been set.

### GetLastFailureReason

`func (o *MsgVpnRestDeliveryPointQueueBinding) GetLastFailureReason() string`

GetLastFailureReason returns the LastFailureReason field if non-nil, zero value otherwise.

### GetLastFailureReasonOk

`func (o *MsgVpnRestDeliveryPointQueueBinding) GetLastFailureReasonOk() (*string, bool)`

GetLastFailureReasonOk returns a tuple with the LastFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailureReason

`func (o *MsgVpnRestDeliveryPointQueueBinding) SetLastFailureReason(v string)`

SetLastFailureReason sets LastFailureReason field to given value.

### HasLastFailureReason

`func (o *MsgVpnRestDeliveryPointQueueBinding) HasLastFailureReason() bool`

HasLastFailureReason returns a boolean if a field has been set.

### GetLastFailureTime

`func (o *MsgVpnRestDeliveryPointQueueBinding) GetLastFailureTime() int32`

GetLastFailureTime returns the LastFailureTime field if non-nil, zero value otherwise.

### GetLastFailureTimeOk

`func (o *MsgVpnRestDeliveryPointQueueBinding) GetLastFailureTimeOk() (*int32, bool)`

GetLastFailureTimeOk returns a tuple with the LastFailureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailureTime

`func (o *MsgVpnRestDeliveryPointQueueBinding) SetLastFailureTime(v int32)`

SetLastFailureTime sets LastFailureTime field to given value.

### HasLastFailureTime

`func (o *MsgVpnRestDeliveryPointQueueBinding) HasLastFailureTime() bool`

HasLastFailureTime returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnRestDeliveryPointQueueBinding) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnRestDeliveryPointQueueBinding) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnRestDeliveryPointQueueBinding) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnRestDeliveryPointQueueBinding) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetPostRequestTarget

`func (o *MsgVpnRestDeliveryPointQueueBinding) GetPostRequestTarget() string`

GetPostRequestTarget returns the PostRequestTarget field if non-nil, zero value otherwise.

### GetPostRequestTargetOk

`func (o *MsgVpnRestDeliveryPointQueueBinding) GetPostRequestTargetOk() (*string, bool)`

GetPostRequestTargetOk returns a tuple with the PostRequestTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostRequestTarget

`func (o *MsgVpnRestDeliveryPointQueueBinding) SetPostRequestTarget(v string)`

SetPostRequestTarget sets PostRequestTarget field to given value.

### HasPostRequestTarget

`func (o *MsgVpnRestDeliveryPointQueueBinding) HasPostRequestTarget() bool`

HasPostRequestTarget returns a boolean if a field has been set.

### GetQueueBindingName

`func (o *MsgVpnRestDeliveryPointQueueBinding) GetQueueBindingName() string`

GetQueueBindingName returns the QueueBindingName field if non-nil, zero value otherwise.

### GetQueueBindingNameOk

`func (o *MsgVpnRestDeliveryPointQueueBinding) GetQueueBindingNameOk() (*string, bool)`

GetQueueBindingNameOk returns a tuple with the QueueBindingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueBindingName

`func (o *MsgVpnRestDeliveryPointQueueBinding) SetQueueBindingName(v string)`

SetQueueBindingName sets QueueBindingName field to given value.

### HasQueueBindingName

`func (o *MsgVpnRestDeliveryPointQueueBinding) HasQueueBindingName() bool`

HasQueueBindingName returns a boolean if a field has been set.

### GetRestDeliveryPointName

`func (o *MsgVpnRestDeliveryPointQueueBinding) GetRestDeliveryPointName() string`

GetRestDeliveryPointName returns the RestDeliveryPointName field if non-nil, zero value otherwise.

### GetRestDeliveryPointNameOk

`func (o *MsgVpnRestDeliveryPointQueueBinding) GetRestDeliveryPointNameOk() (*string, bool)`

GetRestDeliveryPointNameOk returns a tuple with the RestDeliveryPointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestDeliveryPointName

`func (o *MsgVpnRestDeliveryPointQueueBinding) SetRestDeliveryPointName(v string)`

SetRestDeliveryPointName sets RestDeliveryPointName field to given value.

### HasRestDeliveryPointName

`func (o *MsgVpnRestDeliveryPointQueueBinding) HasRestDeliveryPointName() bool`

HasRestDeliveryPointName returns a boolean if a field has been set.

### GetUp

`func (o *MsgVpnRestDeliveryPointQueueBinding) GetUp() bool`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *MsgVpnRestDeliveryPointQueueBinding) GetUpOk() (*bool, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *MsgVpnRestDeliveryPointQueueBinding) SetUp(v bool)`

SetUp sets Up field to given value.

### HasUp

`func (o *MsgVpnRestDeliveryPointQueueBinding) HasUp() bool`

HasUp returns a boolean if a field has been set.

### GetUptime

`func (o *MsgVpnRestDeliveryPointQueueBinding) GetUptime() int64`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *MsgVpnRestDeliveryPointQueueBinding) GetUptimeOk() (*int64, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *MsgVpnRestDeliveryPointQueueBinding) SetUptime(v int64)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *MsgVpnRestDeliveryPointQueueBinding) HasUptime() bool`

HasUptime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


