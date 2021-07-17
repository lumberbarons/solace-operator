# MsgVpnBridgeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**MsgVpnBridge**](MsgVpnBridge.md) |  | [optional] 
**Links** | Pointer to [**MsgVpnBridgeLinks**](MsgVpnBridgeLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnBridgeResponse

`func NewMsgVpnBridgeResponse(meta SempMeta, ) *MsgVpnBridgeResponse`

NewMsgVpnBridgeResponse instantiates a new MsgVpnBridgeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnBridgeResponseWithDefaults

`func NewMsgVpnBridgeResponseWithDefaults() *MsgVpnBridgeResponse`

NewMsgVpnBridgeResponseWithDefaults instantiates a new MsgVpnBridgeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MsgVpnBridgeResponse) GetData() MsgVpnBridge`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnBridgeResponse) GetDataOk() (*MsgVpnBridge, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnBridgeResponse) SetData(v MsgVpnBridge)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnBridgeResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnBridgeResponse) GetLinks() MsgVpnBridgeLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnBridgeResponse) GetLinksOk() (*MsgVpnBridgeLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnBridgeResponse) SetLinks(v MsgVpnBridgeLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnBridgeResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnBridgeResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnBridgeResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnBridgeResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


