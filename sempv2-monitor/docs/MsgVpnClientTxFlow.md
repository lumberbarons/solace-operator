# MsgVpnClientTxFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientName** | Pointer to **string** | The name of the Client. | [optional] 
**EndpointName** | Pointer to **string** | The name of the Queue or Topic Endpoint bound. | [optional] 
**EndpointType** | Pointer to **string** | The type of endpoint bound. The allowed values and their meaning are:  &lt;pre&gt; \&quot;queue\&quot; - The Client is bound to a Queue. \&quot;topic-endpoint\&quot; - The Client is bound to a Topic Endpoint. &lt;/pre&gt;  | [optional] 
**FlowId** | Pointer to **int64** | The identifier (ID) of the flow. | [optional] 
**MsgVpnName** | Pointer to **string** | The name of the Message VPN. | [optional] 

## Methods

### NewMsgVpnClientTxFlow

`func NewMsgVpnClientTxFlow() *MsgVpnClientTxFlow`

NewMsgVpnClientTxFlow instantiates a new MsgVpnClientTxFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnClientTxFlowWithDefaults

`func NewMsgVpnClientTxFlowWithDefaults() *MsgVpnClientTxFlow`

NewMsgVpnClientTxFlowWithDefaults instantiates a new MsgVpnClientTxFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientName

`func (o *MsgVpnClientTxFlow) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *MsgVpnClientTxFlow) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *MsgVpnClientTxFlow) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *MsgVpnClientTxFlow) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetEndpointName

`func (o *MsgVpnClientTxFlow) GetEndpointName() string`

GetEndpointName returns the EndpointName field if non-nil, zero value otherwise.

### GetEndpointNameOk

`func (o *MsgVpnClientTxFlow) GetEndpointNameOk() (*string, bool)`

GetEndpointNameOk returns a tuple with the EndpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointName

`func (o *MsgVpnClientTxFlow) SetEndpointName(v string)`

SetEndpointName sets EndpointName field to given value.

### HasEndpointName

`func (o *MsgVpnClientTxFlow) HasEndpointName() bool`

HasEndpointName returns a boolean if a field has been set.

### GetEndpointType

`func (o *MsgVpnClientTxFlow) GetEndpointType() string`

GetEndpointType returns the EndpointType field if non-nil, zero value otherwise.

### GetEndpointTypeOk

`func (o *MsgVpnClientTxFlow) GetEndpointTypeOk() (*string, bool)`

GetEndpointTypeOk returns a tuple with the EndpointType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointType

`func (o *MsgVpnClientTxFlow) SetEndpointType(v string)`

SetEndpointType sets EndpointType field to given value.

### HasEndpointType

`func (o *MsgVpnClientTxFlow) HasEndpointType() bool`

HasEndpointType returns a boolean if a field has been set.

### GetFlowId

`func (o *MsgVpnClientTxFlow) GetFlowId() int64`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *MsgVpnClientTxFlow) GetFlowIdOk() (*int64, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *MsgVpnClientTxFlow) SetFlowId(v int64)`

SetFlowId sets FlowId field to given value.

### HasFlowId

`func (o *MsgVpnClientTxFlow) HasFlowId() bool`

HasFlowId returns a boolean if a field has been set.

### GetMsgVpnName

`func (o *MsgVpnClientTxFlow) GetMsgVpnName() string`

GetMsgVpnName returns the MsgVpnName field if non-nil, zero value otherwise.

### GetMsgVpnNameOk

`func (o *MsgVpnClientTxFlow) GetMsgVpnNameOk() (*string, bool)`

GetMsgVpnNameOk returns a tuple with the MsgVpnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgVpnName

`func (o *MsgVpnClientTxFlow) SetMsgVpnName(v string)`

SetMsgVpnName sets MsgVpnName field to given value.

### HasMsgVpnName

`func (o *MsgVpnClientTxFlow) HasMsgVpnName() bool`

HasMsgVpnName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


