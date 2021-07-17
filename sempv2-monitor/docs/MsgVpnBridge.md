# MsgVpnBridge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AverageRxByteRate** | Pointer to **int64** | The one minute average of the message rate received from the Bridge, in bytes per second (B/sec). Available since 2.13. | [optional] 
**AverageRxMsgRate** | Pointer to **int64** | The one minute average of the message rate received from the Bridge, in messages per second (msg/sec). Available since 2.13. | [optional] 
**AverageTxByteRate** | Pointer to **int64** | The one minute average of the message rate transmitted to the Bridge, in bytes per second (B/sec). Available since 2.13. | [optional] 
**AverageTxMsgRate** | Pointer to **int64** | The one minute average of the message rate transmitted to the Bridge, in messages per second (msg/sec). Available since 2.13. | [optional] 
**BoundToQueue** | Pointer to **bool** | Indicates whether the Bridge is bound to the queue in the remote Message VPN. | [optional] 
**BridgeName** | Pointer to **string** | The name of the Bridge. | [optional] 
**BridgeVirtualRouter** | Pointer to **string** | The virtual router of the Bridge. The allowed values and their meaning are:  &lt;pre&gt; \&quot;primary\&quot; - The Bridge is used for the primary virtual router. \&quot;backup\&quot; - The Bridge is used for the backup virtual router. \&quot;auto\&quot; - The Bridge is automatically assigned a virtual router at creation, depending on the broker&#39;s active-standby role. &lt;/pre&gt;  | [optional] 
**ClientName** | Pointer to **string** | The name of the Client for the Bridge. | [optional] 
**Compressed** | Pointer to **bool** | Indicates whether messages transmitted over the Bridge are compressed. | [optional] 
**ControlRxByteCount** | Pointer to **int64** | The amount of client control messages received from the Bridge, in bytes (B). Available since 2.13. | [optional] 
**ControlRxMsgCount** | Pointer to **int64** | The number of client control messages received from the Bridge. Available since 2.13. | [optional] 
**ControlTxByteCount** | Pointer to **int64** | The amount of client control messages transmitted to the Bridge, in bytes (B). Available since 2.13. | [optional] 
**ControlTxMsgCount** | Pointer to **int64** | The number of client control messages transmitted to the Bridge. Available since 2.13. | [optional] 
**Counter** | Pointer to [**MsgVpnBridgeCounter**](MsgVpnBridgeCounter.md) |  | [optional] 
**DataRxByteCount** | Pointer to **int64** | The amount of client data messages received from the Bridge, in bytes (B). Available since 2.13. | [optional] 
**DataRxMsgCount** | Pointer to **int64** | The number of client data messages received from the Bridge. Available since 2.13. | [optional] 
**DataTxByteCount** | Pointer to **int64** | The amount of client data messages transmitted to the Bridge, in bytes (B). Available since 2.13. | [optional] 
**DataTxMsgCount** | Pointer to **int64** | The number of client data messages transmitted to the Bridge. Available since 2.13. | [optional] 
**DiscardedRxMsgCount** | Pointer to **int64** | The number of messages discarded during reception from the Bridge. Available since 2.13. | [optional] 
**DiscardedTxMsgCount** | Pointer to **int64** | The number of messages discarded during transmission to the Bridge. Available since 2.13. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the Bridge is enabled. | [optional] 
**Encrypted** | Pointer to **bool** | Indicates whether messages transmitted over the Bridge are encrypted with TLS. | [optional] 
**Establisher** | Pointer to **string** | The establisher of the Bridge connection. The allowed values and their meaning are:  &lt;pre&gt; \&quot;local\&quot; - The Bridge connection was established by the local Message VPN. \&quot;remote\&quot; - The Bridge connection was established by the remote Message VPN. &lt;/pre&gt;  | [optional] 
**InboundFailureReason** | Pointer to **string** | The reason for the inbound connection failure from the Bridge. If there is no failure reason, an empty string (\&quot;\&quot;) is returned. | [optional] 
**InboundState** | Pointer to **string** | The state of the inbound connection from the Bridge. The allowed values and their meaning are:  &lt;pre&gt; \&quot;init\&quot; - The connection is initializing. \&quot;disabled\&quot; - The connection is disabled by configuration. \&quot;enabled\&quot; - The connection is enabled by configuration. \&quot;prepare\&quot; - The connection is operationally down. \&quot;prepare-wait-to-connect\&quot; - The connection is waiting to connect. \&quot;prepare-fetching-dns\&quot; - The domain name of the destination node is being resolved. \&quot;not-ready\&quot; - The connection is operationally down. \&quot;not-ready-connecting\&quot; - The connection is trying to connect. \&quot;not-ready-handshaking\&quot; - The connection is handshaking. \&quot;not-ready-wait-next\&quot; - The connection failed to connect and is waiting to retry. \&quot;not-ready-wait-reuse\&quot; - The connection is closing in order to reuse an existing connection. \&quot;not-ready-wait-bridge-version-mismatch\&quot; - The connection is closing because of a version mismatch. \&quot;not-ready-wait-cleanup\&quot; - The connection is closed and cleaning up. \&quot;ready\&quot; - The connection is operationally up. \&quot;ready-subscribing\&quot; - The connection is up and synchronizing subscriptions. \&quot;ready-in-sync\&quot; - The connection is up and subscriptions are synchronized. &lt;/pre&gt;  | [optional] 
**LastTxMsgId** | Pointer to **int64** | The ID of the last message transmitted to the Bridge. | [optional] 
**LocalInterface** | Pointer to **string** | The physical interface on the local Message VPN host for connecting to the remote Message VPN. | [optional] 
**LocalQueueName** | Pointer to **string** | The name of the local queue for the Bridge. | [optional] 
**LoginRxMsgCount** | Pointer to **int64** | The number of login request messages received from the Bridge. Available since 2.13. | [optional] 
**LoginTxMsgCount** | Pointer to **int64** | The number of login response messages transmitted to the Bridge. Available since 2.13. | [optional] 
**MaxTtl** | Pointer to **int64** | The maximum time-to-live (TTL) in hops. Messages are discarded if their TTL exceeds this value. | [optional] 
**MsgSpoolRxMsgCount** | Pointer to **int64** | The number of guaranteed messages received from the Bridge. Available since 2.13. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**OutboundState** | Pointer to **string** | The state of the outbound connection to the Bridge. The allowed values and their meaning are:  &lt;pre&gt; \&quot;init\&quot; - The connection is initializing. \&quot;disabled\&quot; - The connection is disabled by configuration. \&quot;enabled\&quot; - The connection is enabled by configuration. \&quot;prepare\&quot; - The connection is operationally down. \&quot;prepare-wait-to-connect\&quot; - The connection is waiting to connect. \&quot;prepare-fetching-dns\&quot; - The domain name of the destination node is being resolved. \&quot;not-ready\&quot; - The connection is operationally down. \&quot;not-ready-connecting\&quot; - The connection is trying to connect. \&quot;not-ready-handshaking\&quot; - The connection is handshaking. \&quot;not-ready-wait-next\&quot; - The connection failed to connect and is waiting to retry. \&quot;not-ready-wait-reuse\&quot; - The connection is closing in order to reuse an existing connection. \&quot;not-ready-wait-bridge-version-mismatch\&quot; - The connection is closing because of a version mismatch. \&quot;not-ready-wait-cleanup\&quot; - The connection is closed and cleaning up. \&quot;ready\&quot; - The connection is operationally up. \&quot;ready-subscribing\&quot; - The connection is up and synchronizing subscriptions. \&quot;ready-in-sync\&quot; - The connection is up and subscriptions are synchronized. &lt;/pre&gt;  | [optional] 
**Rate** | Pointer to [**MsgVpnBridgeRate**](MsgVpnBridgeRate.md) |  | [optional] 
**RemoteAddress** | Pointer to **string** | The FQDN or IP address of the remote Message VPN. | [optional] 
**RemoteAuthenticationBasicClientUsername** | Pointer to **string** | The Client Username the Bridge uses to login to the remote Message VPN. | [optional] 
**RemoteAuthenticationScheme** | Pointer to **string** | The authentication scheme for the remote Message VPN. The allowed values and their meaning are:  &lt;pre&gt; \&quot;basic\&quot; - Basic Authentication Scheme (via username and password). \&quot;client-certificate\&quot; - Client Certificate Authentication Scheme (via certificate file or content). &lt;/pre&gt;  | [optional] 
**RemoteConnectionRetryCount** | Pointer to **int64** | The maximum number of retry attempts to establish a connection to the remote Message VPN. A value of 0 means to retry forever. | [optional] 
**RemoteConnectionRetryDelay** | Pointer to **int64** | The number of seconds the broker waits for the bridge connection to be established before attempting a new connection. | [optional] 
**RemoteDeliverToOnePriority** | Pointer to **string** | The priority for deliver-to-one (DTO) messages transmitted from the remote Message VPN. The allowed values and their meaning are:  &lt;pre&gt; \&quot;p1\&quot; - The 1st or highest priority. \&quot;p2\&quot; - The 2nd highest priority. \&quot;p3\&quot; - The 3rd highest priority. \&quot;p4\&quot; - The 4th highest priority. \&quot;da\&quot; - Ignore priority and deliver always. &lt;/pre&gt;  | [optional] 
**RemoteMsgVpnName** | Pointer to **string** | The name of the remote Message VPN. | [optional] 
**RemoteRouterName** | Pointer to **string** | The name of the remote router. | [optional] 
**RemoteTxFlowId** | Pointer to **int32** | The ID of the transmit flow for the connected remote Message VPN. | [optional] 
**RxByteCount** | Pointer to **int64** | The amount of messages received from the Bridge, in bytes (B). Available since 2.13. | [optional] 
**RxByteRate** | Pointer to **int64** | The current message rate received from the Bridge, in bytes per second (B/sec). Available since 2.13. | [optional] 
**RxConnectionFailureCategory** | Pointer to **string** | The category of the inbound connection failure from the Bridge. The allowed values and their meaning are:  &lt;pre&gt; \&quot;no-failure\&quot; - There is no bridge failure. \&quot;local-configuration-problem\&quot; - The bridge failure is a local configuration problem. \&quot;local-operational-state-problem\&quot; - The bridge failure is an operational state problem. &lt;/pre&gt;  Available since 2.18. | [optional] 
**RxMsgCount** | Pointer to **int64** | The number of messages received from the Bridge. Available since 2.13. | [optional] 
**RxMsgRate** | Pointer to **int64** | The current message rate received from the Bridge, in messages per second (msg/sec). Available since 2.13. | [optional] 
**TlsCipherSuiteList** | Pointer to **string** | The colon-separated list of cipher suites supported for TLS connections to the remote Message VPN. The value \&quot;default\&quot; implies all supported suites ordered from most secure to least secure. | [optional] 
**TlsDefaultCipherSuiteList** | Pointer to **bool** | Indicates whether the Bridge is configured to use the default cipher-suite list. | [optional] 
**TtlExceededEventRaised** | Pointer to **bool** | Indicates whether the TTL (hops) exceeded event has been raised. | [optional] 
**TxByteCount** | Pointer to **int64** | The amount of messages transmitted to the Bridge, in bytes (B). Available since 2.13. | [optional] 
**TxByteRate** | Pointer to **int64** | The current message rate transmitted to the Bridge, in bytes per second (B/sec). Available since 2.13. | [optional] 
**TxMsgCount** | Pointer to **int64** | The number of messages transmitted to the Bridge. Available since 2.13. | [optional] 
**TxMsgRate** | Pointer to **int64** | The current message rate transmitted to the Bridge, in messages per second (msg/sec). Available since 2.13. | [optional] 
**Uptime** | Pointer to **int64** | The amount of time in seconds since the Bridge connected to the remote Message VPN. | [optional] 

## Methods

### NewMsgVpnBridge

`func NewMsgVpnBridge() *MsgVpnBridge`

NewMsgVpnBridge instantiates a new MsgVpnBridge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnBridgeWithDefaults

`func NewMsgVpnBridgeWithDefaults() *MsgVpnBridge`

NewMsgVpnBridgeWithDefaults instantiates a new MsgVpnBridge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverageRxByteRate

`func (o *MsgVpnBridge) GetAverageRxByteRate() int64`

GetAverageRxByteRate returns the AverageRxByteRate field if non-nil, zero value otherwise.

### GetAverageRxByteRateOk

`func (o *MsgVpnBridge) GetAverageRxByteRateOk() (*int64, bool)`

GetAverageRxByteRateOk returns a tuple with the AverageRxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRxByteRate

`func (o *MsgVpnBridge) SetAverageRxByteRate(v int64)`

SetAverageRxByteRate sets AverageRxByteRate field to given value.

### HasAverageRxByteRate

`func (o *MsgVpnBridge) HasAverageRxByteRate() bool`

HasAverageRxByteRate returns a boolean if a field has been set.

### GetAverageRxMsgRate

`func (o *MsgVpnBridge) GetAverageRxMsgRate() int64`

GetAverageRxMsgRate returns the AverageRxMsgRate field if non-nil, zero value otherwise.

### GetAverageRxMsgRateOk

`func (o *MsgVpnBridge) GetAverageRxMsgRateOk() (*int64, bool)`

GetAverageRxMsgRateOk returns a tuple with the AverageRxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRxMsgRate

`func (o *MsgVpnBridge) SetAverageRxMsgRate(v int64)`

SetAverageRxMsgRate sets AverageRxMsgRate field to given value.

### HasAverageRxMsgRate

`func (o *MsgVpnBridge) HasAverageRxMsgRate() bool`

HasAverageRxMsgRate returns a boolean if a field has been set.

### GetAverageTxByteRate

`func (o *MsgVpnBridge) GetAverageTxByteRate() int64`

GetAverageTxByteRate returns the AverageTxByteRate field if non-nil, zero value otherwise.

### GetAverageTxByteRateOk

`func (o *MsgVpnBridge) GetAverageTxByteRateOk() (*int64, bool)`

GetAverageTxByteRateOk returns a tuple with the AverageTxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageTxByteRate

`func (o *MsgVpnBridge) SetAverageTxByteRate(v int64)`

SetAverageTxByteRate sets AverageTxByteRate field to given value.

### HasAverageTxByteRate

`func (o *MsgVpnBridge) HasAverageTxByteRate() bool`

HasAverageTxByteRate returns a boolean if a field has been set.

### GetAverageTxMsgRate

`func (o *MsgVpnBridge) GetAverageTxMsgRate() int64`

GetAverageTxMsgRate returns the AverageTxMsgRate field if non-nil, zero value otherwise.

### GetAverageTxMsgRateOk

`func (o *MsgVpnBridge) GetAverageTxMsgRateOk() (*int64, bool)`

GetAverageTxMsgRateOk returns a tuple with the AverageTxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageTxMsgRate

`func (o *MsgVpnBridge) SetAverageTxMsgRate(v int64)`

SetAverageTxMsgRate sets AverageTxMsgRate field to given value.

### HasAverageTxMsgRate

`func (o *MsgVpnBridge) HasAverageTxMsgRate() bool`

HasAverageTxMsgRate returns a boolean if a field has been set.

### GetBoundToQueue

`func (o *MsgVpnBridge) GetBoundToQueue() bool`

GetBoundToQueue returns the BoundToQueue field if non-nil, zero value otherwise.

### GetBoundToQueueOk

`func (o *MsgVpnBridge) GetBoundToQueueOk() (*bool, bool)`

GetBoundToQueueOk returns a tuple with the BoundToQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundToQueue

`func (o *MsgVpnBridge) SetBoundToQueue(v bool)`

SetBoundToQueue sets BoundToQueue field to given value.

### HasBoundToQueue

`func (o *MsgVpnBridge) HasBoundToQueue() bool`

HasBoundToQueue returns a boolean if a field has been set.

### GetBridgeName

`func (o *MsgVpnBridge) GetBridgeName() string`

GetBridgeName returns the BridgeName field if non-nil, zero value otherwise.

### GetBridgeNameOk

`func (o *MsgVpnBridge) GetBridgeNameOk() (*string, bool)`

GetBridgeNameOk returns a tuple with the BridgeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeName

`func (o *MsgVpnBridge) SetBridgeName(v string)`

SetBridgeName sets BridgeName field to given value.

### HasBridgeName

`func (o *MsgVpnBridge) HasBridgeName() bool`

HasBridgeName returns a boolean if a field has been set.

### GetBridgeVirtualRouter

`func (o *MsgVpnBridge) GetBridgeVirtualRouter() string`

GetBridgeVirtualRouter returns the BridgeVirtualRouter field if non-nil, zero value otherwise.

### GetBridgeVirtualRouterOk

`func (o *MsgVpnBridge) GetBridgeVirtualRouterOk() (*string, bool)`

GetBridgeVirtualRouterOk returns a tuple with the BridgeVirtualRouter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeVirtualRouter

`func (o *MsgVpnBridge) SetBridgeVirtualRouter(v string)`

SetBridgeVirtualRouter sets BridgeVirtualRouter field to given value.

### HasBridgeVirtualRouter

`func (o *MsgVpnBridge) HasBridgeVirtualRouter() bool`

HasBridgeVirtualRouter returns a boolean if a field has been set.

### GetClientName

`func (o *MsgVpnBridge) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *MsgVpnBridge) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *MsgVpnBridge) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *MsgVpnBridge) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetCompressed

`func (o *MsgVpnBridge) GetCompressed() bool`

GetCompressed returns the Compressed field if non-nil, zero value otherwise.

### GetCompressedOk

`func (o *MsgVpnBridge) GetCompressedOk() (*bool, bool)`

GetCompressedOk returns a tuple with the Compressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressed

`func (o *MsgVpnBridge) SetCompressed(v bool)`

SetCompressed sets Compressed field to given value.

### HasCompressed

`func (o *MsgVpnBridge) HasCompressed() bool`

HasCompressed returns a boolean if a field has been set.

### GetControlRxByteCount

`func (o *MsgVpnBridge) GetControlRxByteCount() int64`

GetControlRxByteCount returns the ControlRxByteCount field if non-nil, zero value otherwise.

### GetControlRxByteCountOk

`func (o *MsgVpnBridge) GetControlRxByteCountOk() (*int64, bool)`

GetControlRxByteCountOk returns a tuple with the ControlRxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlRxByteCount

`func (o *MsgVpnBridge) SetControlRxByteCount(v int64)`

SetControlRxByteCount sets ControlRxByteCount field to given value.

### HasControlRxByteCount

`func (o *MsgVpnBridge) HasControlRxByteCount() bool`

HasControlRxByteCount returns a boolean if a field has been set.

### GetControlRxMsgCount

`func (o *MsgVpnBridge) GetControlRxMsgCount() int64`

GetControlRxMsgCount returns the ControlRxMsgCount field if non-nil, zero value otherwise.

### GetControlRxMsgCountOk

`func (o *MsgVpnBridge) GetControlRxMsgCountOk() (*int64, bool)`

GetControlRxMsgCountOk returns a tuple with the ControlRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlRxMsgCount

`func (o *MsgVpnBridge) SetControlRxMsgCount(v int64)`

SetControlRxMsgCount sets ControlRxMsgCount field to given value.

### HasControlRxMsgCount

`func (o *MsgVpnBridge) HasControlRxMsgCount() bool`

HasControlRxMsgCount returns a boolean if a field has been set.

### GetControlTxByteCount

`func (o *MsgVpnBridge) GetControlTxByteCount() int64`

GetControlTxByteCount returns the ControlTxByteCount field if non-nil, zero value otherwise.

### GetControlTxByteCountOk

`func (o *MsgVpnBridge) GetControlTxByteCountOk() (*int64, bool)`

GetControlTxByteCountOk returns a tuple with the ControlTxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlTxByteCount

`func (o *MsgVpnBridge) SetControlTxByteCount(v int64)`

SetControlTxByteCount sets ControlTxByteCount field to given value.

### HasControlTxByteCount

`func (o *MsgVpnBridge) HasControlTxByteCount() bool`

HasControlTxByteCount returns a boolean if a field has been set.

### GetControlTxMsgCount

`func (o *MsgVpnBridge) GetControlTxMsgCount() int64`

GetControlTxMsgCount returns the ControlTxMsgCount field if non-nil, zero value otherwise.

### GetControlTxMsgCountOk

`func (o *MsgVpnBridge) GetControlTxMsgCountOk() (*int64, bool)`

GetControlTxMsgCountOk returns a tuple with the ControlTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlTxMsgCount

`func (o *MsgVpnBridge) SetControlTxMsgCount(v int64)`

SetControlTxMsgCount sets ControlTxMsgCount field to given value.

### HasControlTxMsgCount

`func (o *MsgVpnBridge) HasControlTxMsgCount() bool`

HasControlTxMsgCount returns a boolean if a field has been set.

### GetCounter

`func (o *MsgVpnBridge) GetCounter() MsgVpnBridgeCounter`

GetCounter returns the Counter field if non-nil, zero value otherwise.

### GetCounterOk

`func (o *MsgVpnBridge) GetCounterOk() (*MsgVpnBridgeCounter, bool)`

GetCounterOk returns a tuple with the Counter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounter

`func (o *MsgVpnBridge) SetCounter(v MsgVpnBridgeCounter)`

SetCounter sets Counter field to given value.

### HasCounter

`func (o *MsgVpnBridge) HasCounter() bool`

HasCounter returns a boolean if a field has been set.

### GetDataRxByteCount

`func (o *MsgVpnBridge) GetDataRxByteCount() int64`

GetDataRxByteCount returns the DataRxByteCount field if non-nil, zero value otherwise.

### GetDataRxByteCountOk

`func (o *MsgVpnBridge) GetDataRxByteCountOk() (*int64, bool)`

GetDataRxByteCountOk returns a tuple with the DataRxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRxByteCount

`func (o *MsgVpnBridge) SetDataRxByteCount(v int64)`

SetDataRxByteCount sets DataRxByteCount field to given value.

### HasDataRxByteCount

`func (o *MsgVpnBridge) HasDataRxByteCount() bool`

HasDataRxByteCount returns a boolean if a field has been set.

### GetDataRxMsgCount

`func (o *MsgVpnBridge) GetDataRxMsgCount() int64`

GetDataRxMsgCount returns the DataRxMsgCount field if non-nil, zero value otherwise.

### GetDataRxMsgCountOk

`func (o *MsgVpnBridge) GetDataRxMsgCountOk() (*int64, bool)`

GetDataRxMsgCountOk returns a tuple with the DataRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRxMsgCount

`func (o *MsgVpnBridge) SetDataRxMsgCount(v int64)`

SetDataRxMsgCount sets DataRxMsgCount field to given value.

### HasDataRxMsgCount

`func (o *MsgVpnBridge) HasDataRxMsgCount() bool`

HasDataRxMsgCount returns a boolean if a field has been set.

### GetDataTxByteCount

`func (o *MsgVpnBridge) GetDataTxByteCount() int64`

GetDataTxByteCount returns the DataTxByteCount field if non-nil, zero value otherwise.

### GetDataTxByteCountOk

`func (o *MsgVpnBridge) GetDataTxByteCountOk() (*int64, bool)`

GetDataTxByteCountOk returns a tuple with the DataTxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTxByteCount

`func (o *MsgVpnBridge) SetDataTxByteCount(v int64)`

SetDataTxByteCount sets DataTxByteCount field to given value.

### HasDataTxByteCount

`func (o *MsgVpnBridge) HasDataTxByteCount() bool`

HasDataTxByteCount returns a boolean if a field has been set.

### GetDataTxMsgCount

`func (o *MsgVpnBridge) GetDataTxMsgCount() int64`

GetDataTxMsgCount returns the DataTxMsgCount field if non-nil, zero value otherwise.

### GetDataTxMsgCountOk

`func (o *MsgVpnBridge) GetDataTxMsgCountOk() (*int64, bool)`

GetDataTxMsgCountOk returns a tuple with the DataTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTxMsgCount

`func (o *MsgVpnBridge) SetDataTxMsgCount(v int64)`

SetDataTxMsgCount sets DataTxMsgCount field to given value.

### HasDataTxMsgCount

`func (o *MsgVpnBridge) HasDataTxMsgCount() bool`

HasDataTxMsgCount returns a boolean if a field has been set.

### GetDiscardedRxMsgCount

`func (o *MsgVpnBridge) GetDiscardedRxMsgCount() int64`

GetDiscardedRxMsgCount returns the DiscardedRxMsgCount field if non-nil, zero value otherwise.

### GetDiscardedRxMsgCountOk

`func (o *MsgVpnBridge) GetDiscardedRxMsgCountOk() (*int64, bool)`

GetDiscardedRxMsgCountOk returns a tuple with the DiscardedRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscardedRxMsgCount

`func (o *MsgVpnBridge) SetDiscardedRxMsgCount(v int64)`

SetDiscardedRxMsgCount sets DiscardedRxMsgCount field to given value.

### HasDiscardedRxMsgCount

`func (o *MsgVpnBridge) HasDiscardedRxMsgCount() bool`

HasDiscardedRxMsgCount returns a boolean if a field has been set.

### GetDiscardedTxMsgCount

`func (o *MsgVpnBridge) GetDiscardedTxMsgCount() int64`

GetDiscardedTxMsgCount returns the DiscardedTxMsgCount field if non-nil, zero value otherwise.

### GetDiscardedTxMsgCountOk

`func (o *MsgVpnBridge) GetDiscardedTxMsgCountOk() (*int64, bool)`

GetDiscardedTxMsgCountOk returns a tuple with the DiscardedTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscardedTxMsgCount

`func (o *MsgVpnBridge) SetDiscardedTxMsgCount(v int64)`

SetDiscardedTxMsgCount sets DiscardedTxMsgCount field to given value.

### HasDiscardedTxMsgCount

`func (o *MsgVpnBridge) HasDiscardedTxMsgCount() bool`

HasDiscardedTxMsgCount returns a boolean if a field has been set.

### GetEnabled

`func (o *MsgVpnBridge) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MsgVpnBridge) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MsgVpnBridge) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MsgVpnBridge) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEncrypted

`func (o *MsgVpnBridge) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *MsgVpnBridge) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *MsgVpnBridge) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *MsgVpnBridge) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### GetEstablisher

`func (o *MsgVpnBridge) GetEstablisher() string`

GetEstablisher returns the Establisher field if non-nil, zero value otherwise.

### GetEstablisherOk

`func (o *MsgVpnBridge) GetEstablisherOk() (*string, bool)`

GetEstablisherOk returns a tuple with the Establisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstablisher

`func (o *MsgVpnBridge) SetEstablisher(v string)`

SetEstablisher sets Establisher field to given value.

### HasEstablisher

`func (o *MsgVpnBridge) HasEstablisher() bool`

HasEstablisher returns a boolean if a field has been set.

### GetInboundFailureReason

`func (o *MsgVpnBridge) GetInboundFailureReason() string`

GetInboundFailureReason returns the InboundFailureReason field if non-nil, zero value otherwise.

### GetInboundFailureReasonOk

`func (o *MsgVpnBridge) GetInboundFailureReasonOk() (*string, bool)`

GetInboundFailureReasonOk returns a tuple with the InboundFailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundFailureReason

`func (o *MsgVpnBridge) SetInboundFailureReason(v string)`

SetInboundFailureReason sets InboundFailureReason field to given value.

### HasInboundFailureReason

`func (o *MsgVpnBridge) HasInboundFailureReason() bool`

HasInboundFailureReason returns a boolean if a field has been set.

### GetInboundState

`func (o *MsgVpnBridge) GetInboundState() string`

GetInboundState returns the InboundState field if non-nil, zero value otherwise.

### GetInboundStateOk

`func (o *MsgVpnBridge) GetInboundStateOk() (*string, bool)`

GetInboundStateOk returns a tuple with the InboundState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundState

`func (o *MsgVpnBridge) SetInboundState(v string)`

SetInboundState sets InboundState field to given value.

### HasInboundState

`func (o *MsgVpnBridge) HasInboundState() bool`

HasInboundState returns a boolean if a field has been set.

### GetLastTxMsgId

`func (o *MsgVpnBridge) GetLastTxMsgId() int64`

GetLastTxMsgId returns the LastTxMsgId field if non-nil, zero value otherwise.

### GetLastTxMsgIdOk

`func (o *MsgVpnBridge) GetLastTxMsgIdOk() (*int64, bool)`

GetLastTxMsgIdOk returns a tuple with the LastTxMsgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTxMsgId

`func (o *MsgVpnBridge) SetLastTxMsgId(v int64)`

SetLastTxMsgId sets LastTxMsgId field to given value.

### HasLastTxMsgId

`func (o *MsgVpnBridge) HasLastTxMsgId() bool`

HasLastTxMsgId returns a boolean if a field has been set.

### GetLocalInterface

`func (o *MsgVpnBridge) GetLocalInterface() string`

GetLocalInterface returns the LocalInterface field if non-nil, zero value otherwise.

### GetLocalInterfaceOk

`func (o *MsgVpnBridge) GetLocalInterfaceOk() (*string, bool)`

GetLocalInterfaceOk returns a tuple with the LocalInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalInterface

`func (o *MsgVpnBridge) SetLocalInterface(v string)`

SetLocalInterface sets LocalInterface field to given value.

### HasLocalInterface

`func (o *MsgVpnBridge) HasLocalInterface() bool`

HasLocalInterface returns a boolean if a field has been set.

### GetLocalQueueName

`func (o *MsgVpnBridge) GetLocalQueueName() string`

GetLocalQueueName returns the LocalQueueName field if non-nil, zero value otherwise.

### GetLocalQueueNameOk

`func (o *MsgVpnBridge) GetLocalQueueNameOk() (*string, bool)`

GetLocalQueueNameOk returns a tuple with the LocalQueueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalQueueName

`func (o *MsgVpnBridge) SetLocalQueueName(v string)`

SetLocalQueueName sets LocalQueueName field to given value.

### HasLocalQueueName

`func (o *MsgVpnBridge) HasLocalQueueName() bool`

HasLocalQueueName returns a boolean if a field has been set.

### GetLoginRxMsgCount

`func (o *MsgVpnBridge) GetLoginRxMsgCount() int64`

GetLoginRxMsgCount returns the LoginRxMsgCount field if non-nil, zero value otherwise.

### GetLoginRxMsgCountOk

`func (o *MsgVpnBridge) GetLoginRxMsgCountOk() (*int64, bool)`

GetLoginRxMsgCountOk returns a tuple with the LoginRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginRxMsgCount

`func (o *MsgVpnBridge) SetLoginRxMsgCount(v int64)`

SetLoginRxMsgCount sets LoginRxMsgCount field to given value.

### HasLoginRxMsgCount

`func (o *MsgVpnBridge) HasLoginRxMsgCount() bool`

HasLoginRxMsgCount returns a boolean if a field has been set.

### GetLoginTxMsgCount

`func (o *MsgVpnBridge) GetLoginTxMsgCount() int64`

GetLoginTxMsgCount returns the LoginTxMsgCount field if non-nil, zero value otherwise.

### GetLoginTxMsgCountOk

`func (o *MsgVpnBridge) GetLoginTxMsgCountOk() (*int64, bool)`

GetLoginTxMsgCountOk returns a tuple with the LoginTxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginTxMsgCount

`func (o *MsgVpnBridge) SetLoginTxMsgCount(v int64)`

SetLoginTxMsgCount sets LoginTxMsgCount field to given value.

### HasLoginTxMsgCount

`func (o *MsgVpnBridge) HasLoginTxMsgCount() bool`

HasLoginTxMsgCount returns a boolean if a field has been set.

### GetMaxTtl

`func (o *MsgVpnBridge) GetMaxTtl() int64`

GetMaxTtl returns the MaxTtl field if non-nil, zero value otherwise.

### GetMaxTtlOk

`func (o *MsgVpnBridge) GetMaxTtlOk() (*int64, bool)`

GetMaxTtlOk returns a tuple with the MaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTtl

`func (o *MsgVpnBridge) SetMaxTtl(v int64)`

SetMaxTtl sets MaxTtl field to given value.

### HasMaxTtl

`func (o *MsgVpnBridge) HasMaxTtl() bool`

HasMaxTtl returns a boolean if a field has been set.

### GetMsgSpoolRxMsgCount

`func (o *MsgVpnBridge) GetMsgSpoolRxMsgCount() int64`

GetMsgSpoolRxMsgCount returns the MsgSpoolRxMsgCount field if non-nil, zero value otherwise.

### GetMsgSpoolRxMsgCountOk

`func (o *MsgVpnBridge) GetMsgSpoolRxMsgCountOk() (*int64, bool)`

GetMsgSpoolRxMsgCountOk returns a tuple with the MsgSpoolRxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgSpoolRxMsgCount

`func (o *MsgVpnBridge) SetMsgSpoolRxMsgCount(v int64)`

SetMsgSpoolRxMsgCount sets MsgSpoolRxMsgCount field to given value.

### HasMsgSpoolRxMsgCount

`func (o *MsgVpnBridge) HasMsgSpoolRxMsgCount() bool`

HasMsgSpoolRxMsgCount returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnBridge) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnBridge) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnBridge) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnBridge) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetOutboundState

`func (o *MsgVpnBridge) GetOutboundState() string`

GetOutboundState returns the OutboundState field if non-nil, zero value otherwise.

### GetOutboundStateOk

`func (o *MsgVpnBridge) GetOutboundStateOk() (*string, bool)`

GetOutboundStateOk returns a tuple with the OutboundState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundState

`func (o *MsgVpnBridge) SetOutboundState(v string)`

SetOutboundState sets OutboundState field to given value.

### HasOutboundState

`func (o *MsgVpnBridge) HasOutboundState() bool`

HasOutboundState returns a boolean if a field has been set.

### GetRate

`func (o *MsgVpnBridge) GetRate() MsgVpnBridgeRate`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *MsgVpnBridge) GetRateOk() (*MsgVpnBridgeRate, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *MsgVpnBridge) SetRate(v MsgVpnBridgeRate)`

SetRate sets Rate field to given value.

### HasRate

`func (o *MsgVpnBridge) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetRemoteAddress

`func (o *MsgVpnBridge) GetRemoteAddress() string`

GetRemoteAddress returns the RemoteAddress field if non-nil, zero value otherwise.

### GetRemoteAddressOk

`func (o *MsgVpnBridge) GetRemoteAddressOk() (*string, bool)`

GetRemoteAddressOk returns a tuple with the RemoteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddress

`func (o *MsgVpnBridge) SetRemoteAddress(v string)`

SetRemoteAddress sets RemoteAddress field to given value.

### HasRemoteAddress

`func (o *MsgVpnBridge) HasRemoteAddress() bool`

HasRemoteAddress returns a boolean if a field has been set.

### GetRemoteAuthenticationBasicClientUsername

`func (o *MsgVpnBridge) GetRemoteAuthenticationBasicClientUsername() string`

GetRemoteAuthenticationBasicClientUsername returns the RemoteAuthenticationBasicClientUsername field if non-nil, zero value otherwise.

### GetRemoteAuthenticationBasicClientUsernameOk

`func (o *MsgVpnBridge) GetRemoteAuthenticationBasicClientUsernameOk() (*string, bool)`

GetRemoteAuthenticationBasicClientUsernameOk returns a tuple with the RemoteAuthenticationBasicClientUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAuthenticationBasicClientUsername

`func (o *MsgVpnBridge) SetRemoteAuthenticationBasicClientUsername(v string)`

SetRemoteAuthenticationBasicClientUsername sets RemoteAuthenticationBasicClientUsername field to given value.

### HasRemoteAuthenticationBasicClientUsername

`func (o *MsgVpnBridge) HasRemoteAuthenticationBasicClientUsername() bool`

HasRemoteAuthenticationBasicClientUsername returns a boolean if a field has been set.

### GetRemoteAuthenticationScheme

`func (o *MsgVpnBridge) GetRemoteAuthenticationScheme() string`

GetRemoteAuthenticationScheme returns the RemoteAuthenticationScheme field if non-nil, zero value otherwise.

### GetRemoteAuthenticationSchemeOk

`func (o *MsgVpnBridge) GetRemoteAuthenticationSchemeOk() (*string, bool)`

GetRemoteAuthenticationSchemeOk returns a tuple with the RemoteAuthenticationScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAuthenticationScheme

`func (o *MsgVpnBridge) SetRemoteAuthenticationScheme(v string)`

SetRemoteAuthenticationScheme sets RemoteAuthenticationScheme field to given value.

### HasRemoteAuthenticationScheme

`func (o *MsgVpnBridge) HasRemoteAuthenticationScheme() bool`

HasRemoteAuthenticationScheme returns a boolean if a field has been set.

### GetRemoteConnectionRetryCount

`func (o *MsgVpnBridge) GetRemoteConnectionRetryCount() int64`

GetRemoteConnectionRetryCount returns the RemoteConnectionRetryCount field if non-nil, zero value otherwise.

### GetRemoteConnectionRetryCountOk

`func (o *MsgVpnBridge) GetRemoteConnectionRetryCountOk() (*int64, bool)`

GetRemoteConnectionRetryCountOk returns a tuple with the RemoteConnectionRetryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteConnectionRetryCount

`func (o *MsgVpnBridge) SetRemoteConnectionRetryCount(v int64)`

SetRemoteConnectionRetryCount sets RemoteConnectionRetryCount field to given value.

### HasRemoteConnectionRetryCount

`func (o *MsgVpnBridge) HasRemoteConnectionRetryCount() bool`

HasRemoteConnectionRetryCount returns a boolean if a field has been set.

### GetRemoteConnectionRetryDelay

`func (o *MsgVpnBridge) GetRemoteConnectionRetryDelay() int64`

GetRemoteConnectionRetryDelay returns the RemoteConnectionRetryDelay field if non-nil, zero value otherwise.

### GetRemoteConnectionRetryDelayOk

`func (o *MsgVpnBridge) GetRemoteConnectionRetryDelayOk() (*int64, bool)`

GetRemoteConnectionRetryDelayOk returns a tuple with the RemoteConnectionRetryDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteConnectionRetryDelay

`func (o *MsgVpnBridge) SetRemoteConnectionRetryDelay(v int64)`

SetRemoteConnectionRetryDelay sets RemoteConnectionRetryDelay field to given value.

### HasRemoteConnectionRetryDelay

`func (o *MsgVpnBridge) HasRemoteConnectionRetryDelay() bool`

HasRemoteConnectionRetryDelay returns a boolean if a field has been set.

### GetRemoteDeliverToOnePriority

`func (o *MsgVpnBridge) GetRemoteDeliverToOnePriority() string`

GetRemoteDeliverToOnePriority returns the RemoteDeliverToOnePriority field if non-nil, zero value otherwise.

### GetRemoteDeliverToOnePriorityOk

`func (o *MsgVpnBridge) GetRemoteDeliverToOnePriorityOk() (*string, bool)`

GetRemoteDeliverToOnePriorityOk returns a tuple with the RemoteDeliverToOnePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteDeliverToOnePriority

`func (o *MsgVpnBridge) SetRemoteDeliverToOnePriority(v string)`

SetRemoteDeliverToOnePriority sets RemoteDeliverToOnePriority field to given value.

### HasRemoteDeliverToOnePriority

`func (o *MsgVpnBridge) HasRemoteDeliverToOnePriority() bool`

HasRemoteDeliverToOnePriority returns a boolean if a field has been set.

### GetRemoteMsgVpnName

`func (o *MsgVpnBridge) GetRemoteMsgVpnName() string`

GetRemoteMsgVpnName returns the RemoteMsgVpnName field if non-nil, zero value otherwise.

### GetRemoteMsgVpnNameOk

`func (o *MsgVpnBridge) GetRemoteMsgVpnNameOk() (*string, bool)`

GetRemoteMsgVpnNameOk returns a tuple with the RemoteMsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteMsgVpnName

`func (o *MsgVpnBridge) SetRemoteMsgVpnName(v string)`

SetRemoteMsgVpnName sets RemoteMsgVpnName field to given value.

### HasRemoteMsgVpnName

`func (o *MsgVpnBridge) HasRemoteMsgVpnName() bool`

HasRemoteMsgVpnName returns a boolean if a field has been set.

### GetRemoteRouterName

`func (o *MsgVpnBridge) GetRemoteRouterName() string`

GetRemoteRouterName returns the RemoteRouterName field if non-nil, zero value otherwise.

### GetRemoteRouterNameOk

`func (o *MsgVpnBridge) GetRemoteRouterNameOk() (*string, bool)`

GetRemoteRouterNameOk returns a tuple with the RemoteRouterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteRouterName

`func (o *MsgVpnBridge) SetRemoteRouterName(v string)`

SetRemoteRouterName sets RemoteRouterName field to given value.

### HasRemoteRouterName

`func (o *MsgVpnBridge) HasRemoteRouterName() bool`

HasRemoteRouterName returns a boolean if a field has been set.

### GetRemoteTxFlowId

`func (o *MsgVpnBridge) GetRemoteTxFlowId() int32`

GetRemoteTxFlowId returns the RemoteTxFlowId field if non-nil, zero value otherwise.

### GetRemoteTxFlowIdOk

`func (o *MsgVpnBridge) GetRemoteTxFlowIdOk() (*int32, bool)`

GetRemoteTxFlowIdOk returns a tuple with the RemoteTxFlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteTxFlowId

`func (o *MsgVpnBridge) SetRemoteTxFlowId(v int32)`

SetRemoteTxFlowId sets RemoteTxFlowId field to given value.

### HasRemoteTxFlowId

`func (o *MsgVpnBridge) HasRemoteTxFlowId() bool`

HasRemoteTxFlowId returns a boolean if a field has been set.

### GetRxByteCount

`func (o *MsgVpnBridge) GetRxByteCount() int64`

GetRxByteCount returns the RxByteCount field if non-nil, zero value otherwise.

### GetRxByteCountOk

`func (o *MsgVpnBridge) GetRxByteCountOk() (*int64, bool)`

GetRxByteCountOk returns a tuple with the RxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxByteCount

`func (o *MsgVpnBridge) SetRxByteCount(v int64)`

SetRxByteCount sets RxByteCount field to given value.

### HasRxByteCount

`func (o *MsgVpnBridge) HasRxByteCount() bool`

HasRxByteCount returns a boolean if a field has been set.

### GetRxByteRate

`func (o *MsgVpnBridge) GetRxByteRate() int64`

GetRxByteRate returns the RxByteRate field if non-nil, zero value otherwise.

### GetRxByteRateOk

`func (o *MsgVpnBridge) GetRxByteRateOk() (*int64, bool)`

GetRxByteRateOk returns a tuple with the RxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxByteRate

`func (o *MsgVpnBridge) SetRxByteRate(v int64)`

SetRxByteRate sets RxByteRate field to given value.

### HasRxByteRate

`func (o *MsgVpnBridge) HasRxByteRate() bool`

HasRxByteRate returns a boolean if a field has been set.

### GetRxConnectionFailureCategory

`func (o *MsgVpnBridge) GetRxConnectionFailureCategory() string`

GetRxConnectionFailureCategory returns the RxConnectionFailureCategory field if non-nil, zero value otherwise.

### GetRxConnectionFailureCategoryOk

`func (o *MsgVpnBridge) GetRxConnectionFailureCategoryOk() (*string, bool)`

GetRxConnectionFailureCategoryOk returns a tuple with the RxConnectionFailureCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxConnectionFailureCategory

`func (o *MsgVpnBridge) SetRxConnectionFailureCategory(v string)`

SetRxConnectionFailureCategory sets RxConnectionFailureCategory field to given value.

### HasRxConnectionFailureCategory

`func (o *MsgVpnBridge) HasRxConnectionFailureCategory() bool`

HasRxConnectionFailureCategory returns a boolean if a field has been set.

### GetRxMsgCount

`func (o *MsgVpnBridge) GetRxMsgCount() int64`

GetRxMsgCount returns the RxMsgCount field if non-nil, zero value otherwise.

### GetRxMsgCountOk

`func (o *MsgVpnBridge) GetRxMsgCountOk() (*int64, bool)`

GetRxMsgCountOk returns a tuple with the RxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxMsgCount

`func (o *MsgVpnBridge) SetRxMsgCount(v int64)`

SetRxMsgCount sets RxMsgCount field to given value.

### HasRxMsgCount

`func (o *MsgVpnBridge) HasRxMsgCount() bool`

HasRxMsgCount returns a boolean if a field has been set.

### GetRxMsgRate

`func (o *MsgVpnBridge) GetRxMsgRate() int64`

GetRxMsgRate returns the RxMsgRate field if non-nil, zero value otherwise.

### GetRxMsgRateOk

`func (o *MsgVpnBridge) GetRxMsgRateOk() (*int64, bool)`

GetRxMsgRateOk returns a tuple with the RxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxMsgRate

`func (o *MsgVpnBridge) SetRxMsgRate(v int64)`

SetRxMsgRate sets RxMsgRate field to given value.

### HasRxMsgRate

`func (o *MsgVpnBridge) HasRxMsgRate() bool`

HasRxMsgRate returns a boolean if a field has been set.

### GetTlsCipherSuiteList

`func (o *MsgVpnBridge) GetTlsCipherSuiteList() string`

GetTlsCipherSuiteList returns the TlsCipherSuiteList field if non-nil, zero value otherwise.

### GetTlsCipherSuiteListOk

`func (o *MsgVpnBridge) GetTlsCipherSuiteListOk() (*string, bool)`

GetTlsCipherSuiteListOk returns a tuple with the TlsCipherSuiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCipherSuiteList

`func (o *MsgVpnBridge) SetTlsCipherSuiteList(v string)`

SetTlsCipherSuiteList sets TlsCipherSuiteList field to given value.

### HasTlsCipherSuiteList

`func (o *MsgVpnBridge) HasTlsCipherSuiteList() bool`

HasTlsCipherSuiteList returns a boolean if a field has been set.

### GetTlsDefaultCipherSuiteList

`func (o *MsgVpnBridge) GetTlsDefaultCipherSuiteList() bool`

GetTlsDefaultCipherSuiteList returns the TlsDefaultCipherSuiteList field if non-nil, zero value otherwise.

### GetTlsDefaultCipherSuiteListOk

`func (o *MsgVpnBridge) GetTlsDefaultCipherSuiteListOk() (*bool, bool)`

GetTlsDefaultCipherSuiteListOk returns a tuple with the TlsDefaultCipherSuiteList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsDefaultCipherSuiteList

`func (o *MsgVpnBridge) SetTlsDefaultCipherSuiteList(v bool)`

SetTlsDefaultCipherSuiteList sets TlsDefaultCipherSuiteList field to given value.

### HasTlsDefaultCipherSuiteList

`func (o *MsgVpnBridge) HasTlsDefaultCipherSuiteList() bool`

HasTlsDefaultCipherSuiteList returns a boolean if a field has been set.

### GetTtlExceededEventRaised

`func (o *MsgVpnBridge) GetTtlExceededEventRaised() bool`

GetTtlExceededEventRaised returns the TtlExceededEventRaised field if non-nil, zero value otherwise.

### GetTtlExceededEventRaisedOk

`func (o *MsgVpnBridge) GetTtlExceededEventRaisedOk() (*bool, bool)`

GetTtlExceededEventRaisedOk returns a tuple with the TtlExceededEventRaised field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlExceededEventRaised

`func (o *MsgVpnBridge) SetTtlExceededEventRaised(v bool)`

SetTtlExceededEventRaised sets TtlExceededEventRaised field to given value.

### HasTtlExceededEventRaised

`func (o *MsgVpnBridge) HasTtlExceededEventRaised() bool`

HasTtlExceededEventRaised returns a boolean if a field has been set.

### GetTxByteCount

`func (o *MsgVpnBridge) GetTxByteCount() int64`

GetTxByteCount returns the TxByteCount field if non-nil, zero value otherwise.

### GetTxByteCountOk

`func (o *MsgVpnBridge) GetTxByteCountOk() (*int64, bool)`

GetTxByteCountOk returns a tuple with the TxByteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxByteCount

`func (o *MsgVpnBridge) SetTxByteCount(v int64)`

SetTxByteCount sets TxByteCount field to given value.

### HasTxByteCount

`func (o *MsgVpnBridge) HasTxByteCount() bool`

HasTxByteCount returns a boolean if a field has been set.

### GetTxByteRate

`func (o *MsgVpnBridge) GetTxByteRate() int64`

GetTxByteRate returns the TxByteRate field if non-nil, zero value otherwise.

### GetTxByteRateOk

`func (o *MsgVpnBridge) GetTxByteRateOk() (*int64, bool)`

GetTxByteRateOk returns a tuple with the TxByteRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxByteRate

`func (o *MsgVpnBridge) SetTxByteRate(v int64)`

SetTxByteRate sets TxByteRate field to given value.

### HasTxByteRate

`func (o *MsgVpnBridge) HasTxByteRate() bool`

HasTxByteRate returns a boolean if a field has been set.

### GetTxMsgCount

`func (o *MsgVpnBridge) GetTxMsgCount() int64`

GetTxMsgCount returns the TxMsgCount field if non-nil, zero value otherwise.

### GetTxMsgCountOk

`func (o *MsgVpnBridge) GetTxMsgCountOk() (*int64, bool)`

GetTxMsgCountOk returns a tuple with the TxMsgCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxMsgCount

`func (o *MsgVpnBridge) SetTxMsgCount(v int64)`

SetTxMsgCount sets TxMsgCount field to given value.

### HasTxMsgCount

`func (o *MsgVpnBridge) HasTxMsgCount() bool`

HasTxMsgCount returns a boolean if a field has been set.

### GetTxMsgRate

`func (o *MsgVpnBridge) GetTxMsgRate() int64`

GetTxMsgRate returns the TxMsgRate field if non-nil, zero value otherwise.

### GetTxMsgRateOk

`func (o *MsgVpnBridge) GetTxMsgRateOk() (*int64, bool)`

GetTxMsgRateOk returns a tuple with the TxMsgRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxMsgRate

`func (o *MsgVpnBridge) SetTxMsgRate(v int64)`

SetTxMsgRate sets TxMsgRate field to given value.

### HasTxMsgRate

`func (o *MsgVpnBridge) HasTxMsgRate() bool`

HasTxMsgRate returns a boolean if a field has been set.

### GetUptime

`func (o *MsgVpnBridge) GetUptime() int64`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *MsgVpnBridge) GetUptimeOk() (*int64, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *MsgVpnBridge) SetUptime(v int64)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *MsgVpnBridge) HasUptime() bool`

HasUptime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


