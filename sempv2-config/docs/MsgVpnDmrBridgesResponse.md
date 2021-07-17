# MsgVpnDmrBridgesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]MsgVpnDmrBridge**](MsgVpnDmrBridge.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnDmrBridgeLinks**](MsgVpnDmrBridgeLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnDmrBridgesResponse

`func NewMsgVpnDmrBridgesResponse(meta SempMeta, ) *MsgVpnDmrBridgesResponse`

NewMsgVpnDmrBridgesResponse instantiates a new MsgVpnDmrBridgesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnDmrBridgesResponseWithDefaults

`func NewMsgVpnDmrBridgesResponseWithDefaults() *MsgVpnDmrBridgesResponse`

NewMsgVpnDmrBridgesResponseWithDefaults instantiates a new MsgVpnDmrBridgesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MsgVpnDmrBridgesResponse) GetData() []MsgVpnDmrBridge`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnDmrBridgesResponse) GetDataOk() (*[]MsgVpnDmrBridge, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnDmrBridgesResponse) SetData(v []MsgVpnDmrBridge)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnDmrBridgesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnDmrBridgesResponse) GetLinks() []MsgVpnDmrBridgeLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnDmrBridgesResponse) GetLinksOk() (*[]MsgVpnDmrBridgeLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnDmrBridgesResponse) SetLinks(v []MsgVpnDmrBridgeLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnDmrBridgesResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnDmrBridgesResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnDmrBridgesResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnDmrBridgesResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


