# MsgVpnMqttSessionSubscriptionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]MsgVpnMqttSessionSubscription**](MsgVpnMqttSessionSubscription.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnMqttSessionSubscriptionLinks**](MsgVpnMqttSessionSubscriptionLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnMqttSessionSubscriptionsResponse

`func NewMsgVpnMqttSessionSubscriptionsResponse(meta SempMeta, ) *MsgVpnMqttSessionSubscriptionsResponse`

NewMsgVpnMqttSessionSubscriptionsResponse instantiates a new MsgVpnMqttSessionSubscriptionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnMqttSessionSubscriptionsResponseWithDefaults

`func NewMsgVpnMqttSessionSubscriptionsResponseWithDefaults() *MsgVpnMqttSessionSubscriptionsResponse`

NewMsgVpnMqttSessionSubscriptionsResponseWithDefaults instantiates a new MsgVpnMqttSessionSubscriptionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MsgVpnMqttSessionSubscriptionsResponse) GetData() []MsgVpnMqttSessionSubscription`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnMqttSessionSubscriptionsResponse) GetDataOk() (*[]MsgVpnMqttSessionSubscription, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnMqttSessionSubscriptionsResponse) SetData(v []MsgVpnMqttSessionSubscription)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnMqttSessionSubscriptionsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnMqttSessionSubscriptionsResponse) GetLinks() []MsgVpnMqttSessionSubscriptionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnMqttSessionSubscriptionsResponse) GetLinksOk() (*[]MsgVpnMqttSessionSubscriptionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnMqttSessionSubscriptionsResponse) SetLinks(v []MsgVpnMqttSessionSubscriptionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnMqttSessionSubscriptionsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnMqttSessionSubscriptionsResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnMqttSessionSubscriptionsResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnMqttSessionSubscriptionsResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


