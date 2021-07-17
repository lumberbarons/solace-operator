# MsgVpnRestDeliveryPointResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**MsgVpnRestDeliveryPoint**](MsgVpnRestDeliveryPoint.md) |  | [optional] 
**Links** | Pointer to [**MsgVpnRestDeliveryPointLinks**](MsgVpnRestDeliveryPointLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnRestDeliveryPointResponse

`func NewMsgVpnRestDeliveryPointResponse(meta SempMeta, ) *MsgVpnRestDeliveryPointResponse`

NewMsgVpnRestDeliveryPointResponse instantiates a new MsgVpnRestDeliveryPointResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnRestDeliveryPointResponseWithDefaults

`func NewMsgVpnRestDeliveryPointResponseWithDefaults() *MsgVpnRestDeliveryPointResponse`

NewMsgVpnRestDeliveryPointResponseWithDefaults instantiates a new MsgVpnRestDeliveryPointResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MsgVpnRestDeliveryPointResponse) GetData() MsgVpnRestDeliveryPoint`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnRestDeliveryPointResponse) GetDataOk() (*MsgVpnRestDeliveryPoint, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnRestDeliveryPointResponse) SetData(v MsgVpnRestDeliveryPoint)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnRestDeliveryPointResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnRestDeliveryPointResponse) GetLinks() MsgVpnRestDeliveryPointLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnRestDeliveryPointResponse) GetLinksOk() (*MsgVpnRestDeliveryPointLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnRestDeliveryPointResponse) SetLinks(v MsgVpnRestDeliveryPointLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnRestDeliveryPointResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnRestDeliveryPointResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnRestDeliveryPointResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnRestDeliveryPointResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


