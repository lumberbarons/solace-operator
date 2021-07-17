# MsgVpnQueueSubscriptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**MsgVpnQueueSubscription**](MsgVpnQueueSubscription.md) |  | [optional] 
**Links** | Pointer to [**MsgVpnQueueSubscriptionLinks**](MsgVpnQueueSubscriptionLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnQueueSubscriptionResponse

`func NewMsgVpnQueueSubscriptionResponse(meta SempMeta, ) *MsgVpnQueueSubscriptionResponse`

NewMsgVpnQueueSubscriptionResponse instantiates a new MsgVpnQueueSubscriptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnQueueSubscriptionResponseWithDefaults

`func NewMsgVpnQueueSubscriptionResponseWithDefaults() *MsgVpnQueueSubscriptionResponse`

NewMsgVpnQueueSubscriptionResponseWithDefaults instantiates a new MsgVpnQueueSubscriptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MsgVpnQueueSubscriptionResponse) GetData() MsgVpnQueueSubscription`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnQueueSubscriptionResponse) GetDataOk() (*MsgVpnQueueSubscription, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnQueueSubscriptionResponse) SetData(v MsgVpnQueueSubscription)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnQueueSubscriptionResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnQueueSubscriptionResponse) GetLinks() MsgVpnQueueSubscriptionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnQueueSubscriptionResponse) GetLinksOk() (*MsgVpnQueueSubscriptionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnQueueSubscriptionResponse) SetLinks(v MsgVpnQueueSubscriptionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnQueueSubscriptionResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnQueueSubscriptionResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnQueueSubscriptionResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnQueueSubscriptionResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


