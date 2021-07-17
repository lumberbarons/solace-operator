# DmrClusterLinkChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BridgeName** | Pointer to **string** | The name of the Bridge used by the Channel. | [optional] 
**ClientName** | Pointer to **string** | The name of the Client used by the Channel. | [optional] 
**DmrClusterName** | Pointer to **string** | The name of the Cluster. | [optional] 
**Establisher** | Pointer to **bool** | Indicates whether the local node established the Channel. | [optional] 
**FailureReason** | Pointer to **string** | The failure reason for the Channel being down. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 
**QueueName** | Pointer to **string** | The name of the Queue used by the Channel. | [optional] 
**RemoteAddress** | Pointer to **string** | The FQDN or IP address (and optional port) of the remote node. | [optional] 
**RemoteNodeName** | Pointer to **string** | The name of the node at the remote end of the Link. | [optional] 
**Up** | Pointer to **bool** | Indicates whether the Channel is operationally up. | [optional] 
**Uptime** | Pointer to **int64** | The amount of time in seconds since the Channel was up. | [optional] 

## Methods

### NewDmrClusterLinkChannel

`func NewDmrClusterLinkChannel() *DmrClusterLinkChannel`

NewDmrClusterLinkChannel instantiates a new DmrClusterLinkChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDmrClusterLinkChannelWithDefaults

`func NewDmrClusterLinkChannelWithDefaults() *DmrClusterLinkChannel`

NewDmrClusterLinkChannelWithDefaults instantiates a new DmrClusterLinkChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBridgeName

`func (o *DmrClusterLinkChannel) GetBridgeName() string`

GetBridgeName returns the BridgeName field if non-nil, zero value otherwise.

### GetBridgeNameOk

`func (o *DmrClusterLinkChannel) GetBridgeNameOk() (*string, bool)`

GetBridgeNameOk returns a tuple with the BridgeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeName

`func (o *DmrClusterLinkChannel) SetBridgeName(v string)`

SetBridgeName sets BridgeName field to given value.

### HasBridgeName

`func (o *DmrClusterLinkChannel) HasBridgeName() bool`

HasBridgeName returns a boolean if a field has been set.

### GetClientName

`func (o *DmrClusterLinkChannel) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *DmrClusterLinkChannel) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *DmrClusterLinkChannel) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *DmrClusterLinkChannel) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetDmrClusterName

`func (o *DmrClusterLinkChannel) GetDmrClusterName() string`

GetDmrClusterName returns the DmrClusterName field if non-nil, zero value otherwise.

### GetDmrClusterNameOk

`func (o *DmrClusterLinkChannel) GetDmrClusterNameOk() (*string, bool)`

GetDmrClusterNameOk returns a tuple with the DmrClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmrClusterName

`func (o *DmrClusterLinkChannel) SetDmrClusterName(v string)`

SetDmrClusterName sets DmrClusterName field to given value.

### HasDmrClusterName

`func (o *DmrClusterLinkChannel) HasDmrClusterName() bool`

HasDmrClusterName returns a boolean if a field has been set.

### GetEstablisher

`func (o *DmrClusterLinkChannel) GetEstablisher() bool`

GetEstablisher returns the Establisher field if non-nil, zero value otherwise.

### GetEstablisherOk

`func (o *DmrClusterLinkChannel) GetEstablisherOk() (*bool, bool)`

GetEstablisherOk returns a tuple with the Establisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstablisher

`func (o *DmrClusterLinkChannel) SetEstablisher(v bool)`

SetEstablisher sets Establisher field to given value.

### HasEstablisher

`func (o *DmrClusterLinkChannel) HasEstablisher() bool`

HasEstablisher returns a boolean if a field has been set.

### GetFailureReason

`func (o *DmrClusterLinkChannel) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *DmrClusterLinkChannel) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *DmrClusterLinkChannel) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *DmrClusterLinkChannel) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *DmrClusterLinkChannel) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *DmrClusterLinkChannel) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *DmrClusterLinkChannel) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *DmrClusterLinkChannel) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.

### GetQueueName

`func (o *DmrClusterLinkChannel) GetQueueName() string`

GetQueueName returns the QueueName field if non-nil, zero value otherwise.

### GetQueueNameOk

`func (o *DmrClusterLinkChannel) GetQueueNameOk() (*string, bool)`

GetQueueNameOk returns a tuple with the QueueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueName

`func (o *DmrClusterLinkChannel) SetQueueName(v string)`

SetQueueName sets QueueName field to given value.

### HasQueueName

`func (o *DmrClusterLinkChannel) HasQueueName() bool`

HasQueueName returns a boolean if a field has been set.

### GetRemoteAddress

`func (o *DmrClusterLinkChannel) GetRemoteAddress() string`

GetRemoteAddress returns the RemoteAddress field if non-nil, zero value otherwise.

### GetRemoteAddressOk

`func (o *DmrClusterLinkChannel) GetRemoteAddressOk() (*string, bool)`

GetRemoteAddressOk returns a tuple with the RemoteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddress

`func (o *DmrClusterLinkChannel) SetRemoteAddress(v string)`

SetRemoteAddress sets RemoteAddress field to given value.

### HasRemoteAddress

`func (o *DmrClusterLinkChannel) HasRemoteAddress() bool`

HasRemoteAddress returns a boolean if a field has been set.

### GetRemoteNodeName

`func (o *DmrClusterLinkChannel) GetRemoteNodeName() string`

GetRemoteNodeName returns the RemoteNodeName field if non-nil, zero value otherwise.

### GetRemoteNodeNameOk

`func (o *DmrClusterLinkChannel) GetRemoteNodeNameOk() (*string, bool)`

GetRemoteNodeNameOk returns a tuple with the RemoteNodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteNodeName

`func (o *DmrClusterLinkChannel) SetRemoteNodeName(v string)`

SetRemoteNodeName sets RemoteNodeName field to given value.

### HasRemoteNodeName

`func (o *DmrClusterLinkChannel) HasRemoteNodeName() bool`

HasRemoteNodeName returns a boolean if a field has been set.

### GetUp

`func (o *DmrClusterLinkChannel) GetUp() bool`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *DmrClusterLinkChannel) GetUpOk() (*bool, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *DmrClusterLinkChannel) SetUp(v bool)`

SetUp sets Up field to given value.

### HasUp

`func (o *DmrClusterLinkChannel) HasUp() bool`

HasUp returns a boolean if a field has been set.

### GetUptime

`func (o *DmrClusterLinkChannel) GetUptime() int64`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *DmrClusterLinkChannel) GetUptimeOk() (*int64, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *DmrClusterLinkChannel) SetUptime(v int64)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *DmrClusterLinkChannel) HasUptime() bool`

HasUptime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


