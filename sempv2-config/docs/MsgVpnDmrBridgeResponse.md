# MsgVpnDmrBridgeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**MsgVpnDmrBridge**](MsgVpnDmrBridge.md) |  | [optional] 
**Links** | Pointer to [**MsgVpnDmrBridgeLinks**](MsgVpnDmrBridgeLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnDmrBridgeResponse

`func NewMsgVpnDmrBridgeResponse(meta SempMeta, ) *MsgVpnDmrBridgeResponse`

NewMsgVpnDmrBridgeResponse instantiates a new MsgVpnDmrBridgeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnDmrBridgeResponseWithDefaults

`func NewMsgVpnDmrBridgeResponseWithDefaults() *MsgVpnDmrBridgeResponse`

NewMsgVpnDmrBridgeResponseWithDefaults instantiates a new MsgVpnDmrBridgeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MsgVpnDmrBridgeResponse) GetData() MsgVpnDmrBridge`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnDmrBridgeResponse) GetDataOk() (*MsgVpnDmrBridge, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnDmrBridgeResponse) SetData(v MsgVpnDmrBridge)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnDmrBridgeResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnDmrBridgeResponse) GetLinks() MsgVpnDmrBridgeLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnDmrBridgeResponse) GetLinksOk() (*MsgVpnDmrBridgeLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnDmrBridgeResponse) SetLinks(v MsgVpnDmrBridgeLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnDmrBridgeResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnDmrBridgeResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnDmrBridgeResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnDmrBridgeResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


