# MsgVpnJndiConnectionFactoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**MsgVpnJndiConnectionFactory**](MsgVpnJndiConnectionFactory.md) |  | [optional] 
**Links** | Pointer to [**MsgVpnJndiConnectionFactoryLinks**](MsgVpnJndiConnectionFactoryLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnJndiConnectionFactoryResponse

`func NewMsgVpnJndiConnectionFactoryResponse(meta SempMeta, ) *MsgVpnJndiConnectionFactoryResponse`

NewMsgVpnJndiConnectionFactoryResponse instantiates a new MsgVpnJndiConnectionFactoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnJndiConnectionFactoryResponseWithDefaults

`func NewMsgVpnJndiConnectionFactoryResponseWithDefaults() *MsgVpnJndiConnectionFactoryResponse`

NewMsgVpnJndiConnectionFactoryResponseWithDefaults instantiates a new MsgVpnJndiConnectionFactoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MsgVpnJndiConnectionFactoryResponse) GetData() MsgVpnJndiConnectionFactory`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnJndiConnectionFactoryResponse) GetDataOk() (*MsgVpnJndiConnectionFactory, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnJndiConnectionFactoryResponse) SetData(v MsgVpnJndiConnectionFactory)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnJndiConnectionFactoryResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnJndiConnectionFactoryResponse) GetLinks() MsgVpnJndiConnectionFactoryLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnJndiConnectionFactoryResponse) GetLinksOk() (*MsgVpnJndiConnectionFactoryLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnJndiConnectionFactoryResponse) SetLinks(v MsgVpnJndiConnectionFactoryLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnJndiConnectionFactoryResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnJndiConnectionFactoryResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnJndiConnectionFactoryResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnJndiConnectionFactoryResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


