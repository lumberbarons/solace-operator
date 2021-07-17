# MsgVpnRestDeliveryPointRestConsumersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to [**[]MsgVpnRestDeliveryPointRestConsumerCollections**](MsgVpnRestDeliveryPointRestConsumerCollections.md) |  | [optional] 
**Data** | Pointer to [**[]MsgVpnRestDeliveryPointRestConsumer**](MsgVpnRestDeliveryPointRestConsumer.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnRestDeliveryPointRestConsumerLinks**](MsgVpnRestDeliveryPointRestConsumerLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnRestDeliveryPointRestConsumersResponse

`func NewMsgVpnRestDeliveryPointRestConsumersResponse(meta SempMeta, ) *MsgVpnRestDeliveryPointRestConsumersResponse`

NewMsgVpnRestDeliveryPointRestConsumersResponse instantiates a new MsgVpnRestDeliveryPointRestConsumersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnRestDeliveryPointRestConsumersResponseWithDefaults

`func NewMsgVpnRestDeliveryPointRestConsumersResponseWithDefaults() *MsgVpnRestDeliveryPointRestConsumersResponse`

NewMsgVpnRestDeliveryPointRestConsumersResponseWithDefaults instantiates a new MsgVpnRestDeliveryPointRestConsumersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnRestDeliveryPointRestConsumersResponse) GetCollections() []MsgVpnRestDeliveryPointRestConsumerCollections`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnRestDeliveryPointRestConsumersResponse) GetCollectionsOk() (*[]MsgVpnRestDeliveryPointRestConsumerCollections, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnRestDeliveryPointRestConsumersResponse) SetCollections(v []MsgVpnRestDeliveryPointRestConsumerCollections)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnRestDeliveryPointRestConsumersResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnRestDeliveryPointRestConsumersResponse) GetData() []MsgVpnRestDeliveryPointRestConsumer`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnRestDeliveryPointRestConsumersResponse) GetDataOk() (*[]MsgVpnRestDeliveryPointRestConsumer, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnRestDeliveryPointRestConsumersResponse) SetData(v []MsgVpnRestDeliveryPointRestConsumer)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnRestDeliveryPointRestConsumersResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnRestDeliveryPointRestConsumersResponse) GetLinks() []MsgVpnRestDeliveryPointRestConsumerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnRestDeliveryPointRestConsumersResponse) GetLinksOk() (*[]MsgVpnRestDeliveryPointRestConsumerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnRestDeliveryPointRestConsumersResponse) SetLinks(v []MsgVpnRestDeliveryPointRestConsumerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnRestDeliveryPointRestConsumersResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnRestDeliveryPointRestConsumersResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnRestDeliveryPointRestConsumersResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnRestDeliveryPointRestConsumersResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


