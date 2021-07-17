# MsgVpnBridgeRemoteMsgVpn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoundToQueue** | Pointer to **bool** | Indicates whether the Bridge is bound to the queue in the remote Message VPN. | [optional] 
**BridgeName** | Pointer to **string** | The name of the Bridge. | [optional] 
**BridgeVirtualRouter** | Pointer to **string** | The virtual router of the Bridge. The allowed values and their meaning are:  &lt;pre&gt; \&quot;primary\&quot; - The Bridge is used for the primary virtual router. \&quot;backup\&quot; - The Bridge is used for the backup virtual router. \&quot;auto\&quot; - The Bridge is automatically assigned a virtual router at creation, depending on the broker&#39;s active-standby role. &lt;/pre&gt;  | [optional] 
**ClientUsername** | Pointer to **string** | The Client Username the Bridge uses to login to the remote Message VPN. This per remote Message VPN value overrides the value provided for the Bridge overall. | [optional] 
**CompressedDataEnabled** | Pointer to **bool** | Indicates whether data compression is enabled for the remote Message VPN connection. | [optional] 
**ConnectOrder** | Pointer to **int32** | The preference given to incoming connections from remote Message VPN hosts, from 1 (highest priority) to 4 (lowest priority). | [optional] 
**EgressFlowWindowSize** | Pointer to **int64** | The number of outstanding guaranteed messages that can be transmitted over the remote Message VPN connection before an acknowledgement is received. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the remote Message VPN is enabled. | [optional] 
**LastConnectionFailureReason** | Pointer to **string** | The reason for the last connection failure to the remote Message VPN. | [optional] 
**LastQueueBindFailureReason** | Pointer to **string** | The reason for the last queue bind failure in the remote Message VPN. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**QueueBinding** | Pointer to **string** | The queue binding of the Bridge in the remote Message VPN. | [optional] 
**QueueBoundUptime** | Pointer to **int32** | The amount of time in seconds since the queue was bound in the remote Message VPN. | [optional] 
**RemoteMsgVpnInterface** | Pointer to **string** | The physical interface on the local Message VPN host for connecting to the remote Message VPN. By default, an interface is chosen automatically (recommended), but if specified, &#x60;remoteMsgVpnLocation&#x60; must not be a virtual router name. | [optional] 
**RemoteMsgVpnLocation** | Pointer to **string** | The location of the remote Message VPN as either an FQDN with port, IP address with port, or virtual router name (starting with \&quot;v:\&quot;). | [optional] 
**RemoteMsgVpnName** | Pointer to **string** | The name of the remote Message VPN. | [optional] 
**TlsEnabled** | Pointer to **bool** | Indicates whether encryption (TLS) is enabled for the remote Message VPN connection. | [optional] 
**UnidirectionalClientProfile** | Pointer to **string** | The Client Profile for the unidirectional Bridge of the remote Message VPN. The Client Profile must exist in the local Message VPN, and it is used only for the TCP parameters. Note that the default client profile has a TCP maximum window size of 2MB. | [optional] 
**Up** | Pointer to **bool** | Indicates whether the connection to the remote Message VPN is up. | [optional] 

## Methods

### NewMsgVpnBridgeRemoteMsgVpn

`func NewMsgVpnBridgeRemoteMsgVpn() *MsgVpnBridgeRemoteMsgVpn`

NewMsgVpnBridgeRemoteMsgVpn instantiates a new MsgVpnBridgeRemoteMsgVpn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnBridgeRemoteMsgVpnWithDefaults

`func NewMsgVpnBridgeRemoteMsgVpnWithDefaults() *MsgVpnBridgeRemoteMsgVpn`

NewMsgVpnBridgeRemoteMsgVpnWithDefaults instantiates a new MsgVpnBridgeRemoteMsgVpn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoundToQueue

`func (o *MsgVpnBridgeRemoteMsgVpn) GetBoundToQueue() bool`

GetBoundToQueue returns the BoundToQueue field if non-nil, zero value otherwise.

### GetBoundToQueueOk

`func (o *MsgVpnBridgeRemoteMsgVpn) GetBoundToQueueOk() (*bool, bool)`

GetBoundToQueueOk returns a tuple with the BoundToQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundToQueue

`func (o *MsgVpnBridgeRemoteMsgVpn) SetBoundToQueue(v bool)`

SetBoundToQueue sets BoundToQueue field to given value.

### HasBoundToQueue

`func (o *MsgVpnBridgeRemoteMsgVpn) HasBoundToQueue() bool`

HasBoundToQueue returns a boolean if a field has been set.

### GetBridgeName

`func (o *MsgVpnBridgeRemoteMsgVpn) GetBridgeName() string`

GetBridgeName returns the BridgeName field if non-nil, zero value otherwise.

### GetBridgeNameOk

`func (o *MsgVpnBridgeRemoteMsgVpn) GetBridgeNameOk() (*string, bool)`

GetBridgeNameOk returns a tuple with the BridgeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeName

`func (o *MsgVpnBridgeRemoteMsgVpn) SetBridgeName(v string)`

SetBridgeName sets BridgeName field to given value.

### HasBridgeName

`func (o *MsgVpnBridgeRemoteMsgVpn) HasBridgeName() bool`

HasBridgeName returns a boolean if a field has been set.

### GetBridgeVirtualRouter

`func (o *MsgVpnBridgeRemoteMsgVpn) GetBridgeVirtualRouter() string`

GetBridgeVirtualRouter returns the BridgeVirtualRouter field if non-nil, zero value otherwise.

### GetBridgeVirtualRouterOk

`func (o *MsgVpnBridgeRemoteMsgVpn) GetBridgeVirtualRouterOk() (*string, bool)`

GetBridgeVirtualRouterOk returns a tuple with the BridgeVirtualRouter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeVirtualRouter

`func (o *MsgVpnBridgeRemoteMsgVpn) SetBridgeVirtualRouter(v string)`

SetBridgeVirtualRouter sets BridgeVirtualRouter field to given value.

### HasBridgeVirtualRouter

`func (o *MsgVpnBridgeRemoteMsgVpn) HasBridgeVirtualRouter() bool`

HasBridgeVirtualRouter returns a boolean if a field has been set.

### GetClientUsername

`func (o *MsgVpnBridgeRemoteMsgVpn) GetClientUsername() string`

GetClientUsername returns the ClientUsername field if non-nil, zero value otherwise.

### GetClientUsernameOk

`func (o *MsgVpnBridgeRemoteMsgVpn) GetClientUsernameOk() (*string, bool)`

GetClientUsernameOk returns a tuple with the ClientUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientUsername

`func (o *MsgVpnBridgeRemoteMsgVpn) SetClientUsername(v string)`

SetClientUsername sets ClientUsername field to given value.

### HasClientUsername

`func (o *MsgVpnBridgeRemoteMsgVpn) HasClientUsername() bool`

HasClientUsername returns a boolean if a field has been set.

### GetCompressedDataEnabled

`func (o *MsgVpnBridgeRemoteMsgVpn) GetCompressedDataEnabled() bool`

GetCompressedDataEnabled returns the CompressedDataEnabled field if non-nil, zero value otherwise.

### GetCompressedDataEnabledOk

`func (o *MsgVpnBridgeRemoteMsgVpn) GetCompressedDataEnabledOk() (*bool, bool)`

GetCompressedDataEnabledOk returns a tuple with the CompressedDataEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressedDataEnabled

`func (o *MsgVpnBridgeRemoteMsgVpn) SetCompressedDataEnabled(v bool)`

SetCompressedDataEnabled sets CompressedDataEnabled field to given value.

### HasCompressedDataEnabled

`func (o *MsgVpnBridgeRemoteMsgVpn) HasCompressedDataEnabled() bool`

HasCompressedDataEnabled returns a boolean if a field has been set.

### GetConnectOrder

`func (o *MsgVpnBridgeRemoteMsgVpn) GetConnectOrder() int32`

GetConnectOrder returns the ConnectOrder field if non-nil, zero value otherwise.

### GetConnectOrderOk

`func (o *MsgVpnBridgeRemoteMsgVpn) GetConnectOrderOk() (*int32, bool)`

GetConnectOrderOk returns a tuple with the ConnectOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectOrder

`func (o *MsgVpnBridgeRemoteMsgVpn) SetConnectOrder(v int32)`

SetConnectOrder sets ConnectOrder field to given value.

### HasConnectOrder

`func (o *MsgVpnBridgeRemoteMsgVpn) HasConnectOrder() bool`

HasConnectOrder returns a boolean if a field has been set.

### GetEgressFlowWindowSize

`func (o *MsgVpnBridgeRemoteMsgVpn) GetEgressFlowWindowSize() int64`

GetEgressFlowWindowSize returns the EgressFlowWindowSize field if non-nil, zero value otherwise.

### GetEgressFlowWindowSizeOk

`func (o *MsgVpnBridgeRemoteMsgVpn) GetEgressFlowWindowSizeOk() (*int64, bool)`

GetEgressFlowWindowSizeOk returns a tuple with the EgressFlowWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressFlowWindowSize

`func (o *MsgVpnBridgeRemoteMsgVpn) SetEgressFlowWindowSize(v int64)`

SetEgressFlowWindowSize sets EgressFlowWindowSize field to given value.

### HasEgressFlowWindowSize

`func (o *MsgVpnBridgeRemoteMsgVpn) HasEgressFlowWindowSize() bool`

HasEgressFlowWindowSize returns a boolean if a field has been set.

### GetEnabled

`func (o *MsgVpnBridgeRemoteMsgVpn) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MsgVpnBridgeRemoteMsgVpn) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MsgVpnBridgeRemoteMsgVpn) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MsgVpnBridgeRemoteMsgVpn) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLastConnectionFailureReason

`func (o *MsgVpnBridgeRemoteMsgVpn) GetLastConnectionFailureReason() string`

GetLastConnectionFailureReason returns the LastConnectionFailureReason field if non-nil, zero value otherwise.

### GetLastConnectionFailureReasonOk

`func (o *MsgVpnBridgeRemoteMsgVpn) GetLastConnectionFailureReasonOk() (*string, bool)`

GetLastConnectionFailureReasonOk returns a tuple with the LastConnectionFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastConnectionFailureReason

`func (o *MsgVpnBridgeRemoteMsgVpn) SetLastConnectionFailureReason(v string)`

SetLastConnectionFailureReason sets LastConnectionFailureReason field to given value.

### HasLastConnectionFailureReason

`func (o *MsgVpnBridgeRemoteMsgVpn) HasLastConnectionFailureReason() bool`

HasLastConnectionFailureReason returns a boolean if a field has been set.

### GetLastQueueBindFailureReason

`func (o *MsgVpnBridgeRemoteMsgVpn) GetLastQueueBindFailureReason() string`

GetLastQueueBindFailureReason returns the LastQueueBindFailureReason field if non-nil, zero value otherwise.

### GetLastQueueBindFailureReasonOk

`func (o *MsgVpnBridgeRemoteMsgVpn) GetLastQueueBindFailureReasonOk() (*string, bool)`

GetLastQueueBindFailureReasonOk returns a tuple with the LastQueueBindFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueueBindFailureReason

`func (o *MsgVpnBridgeRemoteMsgVpn) SetLastQueueBindFailureReason(v string)`

SetLastQueueBindFailureReason sets LastQueueBindFailureReason field to given value.

### HasLastQueueBindFailureReason

`func (o *MsgVpnBridgeRemoteMsgVpn) HasLastQueueBindFailureReason() bool`

HasLastQueueBindFailureReason returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnBridgeRemoteMsgVpn) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnBridgeRemoteMsgVpn) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnBridgeRemoteMsgVpn) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnBridgeRemoteMsgVpn) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetQueueBinding

`func (o *MsgVpnBridgeRemoteMsgVpn) GetQueueBinding() string`

GetQueueBinding returns the QueueBinding field if non-nil, zero value otherwise.

### GetQueueBindingOk

`func (o *MsgVpnBridgeRemoteMsgVpn) GetQueueBindingOk() (*string, bool)`

GetQueueBindingOk returns a tuple with the QueueBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueBinding

`func (o *MsgVpnBridgeRemoteMsgVpn) SetQueueBinding(v string)`

SetQueueBinding sets QueueBinding field to given value.

### HasQueueBinding

`func (o *MsgVpnBridgeRemoteMsgVpn) HasQueueBinding() bool`

HasQueueBinding returns a boolean if a field has been set.

### GetQueueBoundUptime

`func (o *MsgVpnBridgeRemoteMsgVpn) GetQueueBoundUptime() int32`

GetQueueBoundUptime returns the QueueBoundUptime field if non-nil, zero value otherwise.

### GetQueueBoundUptimeOk

`func (o *MsgVpnBridgeRemoteMsgVpn) GetQueueBoundUptimeOk() (*int32, bool)`

GetQueueBoundUptimeOk returns a tuple with the QueueBoundUptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueBoundUptime

`func (o *MsgVpnBridgeRemoteMsgVpn) SetQueueBoundUptime(v int32)`

SetQueueBoundUptime sets QueueBoundUptime field to given value.

### HasQueueBoundUptime

`func (o *MsgVpnBridgeRemoteMsgVpn) HasQueueBoundUptime() bool`

HasQueueBoundUptime returns a boolean if a field has been set.

### GetRemoteMsgVpnInterface

`func (o *MsgVpnBridgeRemoteMsgVpn) GetRemoteMsgVpnInterface() string`

GetRemoteMsgVpnInterface returns the RemoteMsgVpnInterface field if non-nil, zero value otherwise.

### GetRemoteMsgVpnInterfaceOk

`func (o *MsgVpnBridgeRemoteMsgVpn) GetRemoteMsgVpnInterfaceOk() (*string, bool)`

GetRemoteMsgVpnInterfaceOk returns a tuple with the RemoteMsgVpnInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteMsgVpnInterface

`func (o *MsgVpnBridgeRemoteMsgVpn) SetRemoteMsgVpnInterface(v string)`

SetRemoteMsgVpnInterface sets RemoteMsgVpnInterface field to given value.

### HasRemoteMsgVpnInterface

`func (o *MsgVpnBridgeRemoteMsgVpn) HasRemoteMsgVpnInterface() bool`

HasRemoteMsgVpnInterface returns a boolean if a field has been set.

### GetRemoteMsgVpnLocation

`func (o *MsgVpnBridgeRemoteMsgVpn) GetRemoteMsgVpnLocation() string`

GetRemoteMsgVpnLocation returns the RemoteMsgVpnLocation field if non-nil, zero value otherwise.

### GetRemoteMsgVpnLocationOk

`func (o *MsgVpnBridgeRemoteMsgVpn) GetRemoteMsgVpnLocationOk() (*string, bool)`

GetRemoteMsgVpnLocationOk returns a tuple with the RemoteMsgVpnLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteMsgVpnLocation

`func (o *MsgVpnBridgeRemoteMsgVpn) SetRemoteMsgVpnLocation(v string)`

SetRemoteMsgVpnLocation sets RemoteMsgVpnLocation field to given value.

### HasRemoteMsgVpnLocation

`func (o *MsgVpnBridgeRemoteMsgVpn) HasRemoteMsgVpnLocation() bool`

HasRemoteMsgVpnLocation returns a boolean if a field has been set.

### GetRemoteMsgVpnName

`func (o *MsgVpnBridgeRemoteMsgVpn) GetRemoteMsgVpnName() string`

GetRemoteMsgVpnName returns the RemoteMsgVpnName field if non-nil, zero value otherwise.

### GetRemoteMsgVpnNameOk

`func (o *MsgVpnBridgeRemoteMsgVpn) GetRemoteMsgVpnNameOk() (*string, bool)`

GetRemoteMsgVpnNameOk returns a tuple with the RemoteMsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteMsgVpnName

`func (o *MsgVpnBridgeRemoteMsgVpn) SetRemoteMsgVpnName(v string)`

SetRemoteMsgVpnName sets RemoteMsgVpnName field to given value.

### HasRemoteMsgVpnName

`func (o *MsgVpnBridgeRemoteMsgVpn) HasRemoteMsgVpnName() bool`

HasRemoteMsgVpnName returns a boolean if a field has been set.

### GetTlsEnabled

`func (o *MsgVpnBridgeRemoteMsgVpn) GetTlsEnabled() bool`

GetTlsEnabled returns the TlsEnabled field if non-nil, zero value otherwise.

### GetTlsEnabledOk

`func (o *MsgVpnBridgeRemoteMsgVpn) GetTlsEnabledOk() (*bool, bool)`

GetTlsEnabledOk returns a tuple with the TlsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsEnabled

`func (o *MsgVpnBridgeRemoteMsgVpn) SetTlsEnabled(v bool)`

SetTlsEnabled sets TlsEnabled field to given value.

### HasTlsEnabled

`func (o *MsgVpnBridgeRemoteMsgVpn) HasTlsEnabled() bool`

HasTlsEnabled returns a boolean if a field has been set.

### GetUnidirectionalClientProfile

`func (o *MsgVpnBridgeRemoteMsgVpn) GetUnidirectionalClientProfile() string`

GetUnidirectionalClientProfile returns the UnidirectionalClientProfile field if non-nil, zero value otherwise.

### GetUnidirectionalClientProfileOk

`func (o *MsgVpnBridgeRemoteMsgVpn) GetUnidirectionalClientProfileOk() (*string, bool)`

GetUnidirectionalClientProfileOk returns a tuple with the UnidirectionalClientProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnidirectionalClientProfile

`func (o *MsgVpnBridgeRemoteMsgVpn) SetUnidirectionalClientProfile(v string)`

SetUnidirectionalClientProfile sets UnidirectionalClientProfile field to given value.

### HasUnidirectionalClientProfile

`func (o *MsgVpnBridgeRemoteMsgVpn) HasUnidirectionalClientProfile() bool`

HasUnidirectionalClientProfile returns a boolean if a field has been set.

### GetUp

`func (o *MsgVpnBridgeRemoteMsgVpn) GetUp() bool`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *MsgVpnBridgeRemoteMsgVpn) GetUpOk() (*bool, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *MsgVpnBridgeRemoteMsgVpn) SetUp(v bool)`

SetUp sets Up field to given value.

### HasUp

`func (o *MsgVpnBridgeRemoteMsgVpn) HasUp() bool`

HasUp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


