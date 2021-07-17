# MsgVpnRestDeliveryPointQueueBindingsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**[]MsgVpnRestDeliveryPointQueueBinding**](MsgVpnRestDeliveryPointQueueBinding.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnRestDeliveryPointQueueBindingLinks**](MsgVpnRestDeliveryPointQueueBindingLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnRestDeliveryPointQueueBindingsResponse

`func NewMsgVpnRestDeliveryPointQueueBindingsResponse(meta SempMeta, ) *MsgVpnRestDeliveryPointQueueBindingsResponse`

NewMsgVpnRestDeliveryPointQueueBindingsResponse instantiates a new MsgVpnRestDeliveryPointQueueBindingsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnRestDeliveryPointQueueBindingsResponseWithDefaults

`func NewMsgVpnRestDeliveryPointQueueBindingsResponseWithDefaults() *MsgVpnRestDeliveryPointQueueBindingsResponse`

NewMsgVpnRestDeliveryPointQueueBindingsResponseWithDefaults instantiates a new MsgVpnRestDeliveryPointQueueBindingsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnRestDeliveryPointQueueBindingsResponse) GetCollections() []map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnRestDeliveryPointQueueBindingsResponse) GetCollectionsOk() (*[]map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnRestDeliveryPointQueueBindingsResponse) SetCollections(v []map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnRestDeliveryPointQueueBindingsResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnRestDeliveryPointQueueBindingsResponse) GetData() []MsgVpnRestDeliveryPointQueueBinding`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnRestDeliveryPointQueueBindingsResponse) GetDataOk() (*[]MsgVpnRestDeliveryPointQueueBinding, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnRestDeliveryPointQueueBindingsResponse) SetData(v []MsgVpnRestDeliveryPointQueueBinding)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnRestDeliveryPointQueueBindingsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnRestDeliveryPointQueueBindingsResponse) GetLinks() []MsgVpnRestDeliveryPointQueueBindingLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnRestDeliveryPointQueueBindingsResponse) GetLinksOk() (*[]MsgVpnRestDeliveryPointQueueBindingLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnRestDeliveryPointQueueBindingsResponse) SetLinks(v []MsgVpnRestDeliveryPointQueueBindingLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnRestDeliveryPointQueueBindingsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnRestDeliveryPointQueueBindingsResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnRestDeliveryPointQueueBindingsResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnRestDeliveryPointQueueBindingsResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


