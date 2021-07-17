# MsgVpnDmrBridge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailureReason** | Pointer to **string** | The last failure reason for the DMR Bridge. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**RemoteMsgVpnName** | Pointer to **string** | The remote Message VPN of the DMR Bridge. | [optional] 
**RemoteNodeName** | Pointer to **string** | The name of the node at the remote end of the DMR Bridge. | [optional] 
**Up** | Pointer to **bool** | Indicates whether the operational state of the DMR Bridge is up. | [optional] 
**Uptime** | Pointer to **int64** | The amount of time in seconds since the DMR Bridge was up. | [optional] 

## Methods

### NewMsgVpnDmrBridge

`func NewMsgVpnDmrBridge() *MsgVpnDmrBridge`

NewMsgVpnDmrBridge instantiates a new MsgVpnDmrBridge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnDmrBridgeWithDefaults

`func NewMsgVpnDmrBridgeWithDefaults() *MsgVpnDmrBridge`

NewMsgVpnDmrBridgeWithDefaults instantiates a new MsgVpnDmrBridge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailureReason

`func (o *MsgVpnDmrBridge) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *MsgVpnDmrBridge) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *MsgVpnDmrBridge) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *MsgVpnDmrBridge) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnDmrBridge) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnDmrBridge) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnDmrBridge) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnDmrBridge) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetRemoteMsgVpnName

`func (o *MsgVpnDmrBridge) GetRemoteMsgVpnName() string`

GetRemoteMsgVpnName returns the RemoteMsgVpnName field if non-nil, zero value otherwise.

### GetRemoteMsgVpnNameOk

`func (o *MsgVpnDmrBridge) GetRemoteMsgVpnNameOk() (*string, bool)`

GetRemoteMsgVpnNameOk returns a tuple with the RemoteMsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteMsgVpnName

`func (o *MsgVpnDmrBridge) SetRemoteMsgVpnName(v string)`

SetRemoteMsgVpnName sets RemoteMsgVpnName field to given value.

### HasRemoteMsgVpnName

`func (o *MsgVpnDmrBridge) HasRemoteMsgVpnName() bool`

HasRemoteMsgVpnName returns a boolean if a field has been set.

### GetRemoteNodeName

`func (o *MsgVpnDmrBridge) GetRemoteNodeName() string`

GetRemoteNodeName returns the RemoteNodeName field if non-nil, zero value otherwise.

### GetRemoteNodeNameOk

`func (o *MsgVpnDmrBridge) GetRemoteNodeNameOk() (*string, bool)`

GetRemoteNodeNameOk returns a tuple with the RemoteNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteNodeName

`func (o *MsgVpnDmrBridge) SetRemoteNodeName(v string)`

SetRemoteNodeName sets RemoteNodeName field to given value.

### HasRemoteNodeName

`func (o *MsgVpnDmrBridge) HasRemoteNodeName() bool`

HasRemoteNodeName returns a boolean if a field has been set.

### GetUp

`func (o *MsgVpnDmrBridge) GetUp() bool`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *MsgVpnDmrBridge) GetUpOk() (*bool, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *MsgVpnDmrBridge) SetUp(v bool)`

SetUp sets Up field to given value.

### HasUp

`func (o *MsgVpnDmrBridge) HasUp() bool`

HasUp returns a boolean if a field has been set.

### GetUptime

`func (o *MsgVpnDmrBridge) GetUptime() int64`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *MsgVpnDmrBridge) GetUptimeOk() (*int64, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *MsgVpnDmrBridge) SetUptime(v int64)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *MsgVpnDmrBridge) HasUptime() bool`

HasUptime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


