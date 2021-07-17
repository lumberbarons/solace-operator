# MsgVpnMqttSessionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to [**MsgVpnMqttSessionCollections**](MsgVpnMqttSessionCollections.md) |  | [optional] 
**Data** | Pointer to [**MsgVpnMqttSession**](MsgVpnMqttSession.md) |  | [optional] 
**Links** | Pointer to [**MsgVpnMqttSessionLinks**](MsgVpnMqttSessionLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnMqttSessionResponse

`func NewMsgVpnMqttSessionResponse(meta SempMeta, ) *MsgVpnMqttSessionResponse`

NewMsgVpnMqttSessionResponse instantiates a new MsgVpnMqttSessionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnMqttSessionResponseWithDefaults

`func NewMsgVpnMqttSessionResponseWithDefaults() *MsgVpnMqttSessionResponse`

NewMsgVpnMqttSessionResponseWithDefaults instantiates a new MsgVpnMqttSessionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnMqttSessionResponse) GetCollections() MsgVpnMqttSessionCollections`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnMqttSessionResponse) GetCollectionsOk() (*MsgVpnMqttSessionCollections, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnMqttSessionResponse) SetCollections(v MsgVpnMqttSessionCollections)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnMqttSessionResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnMqttSessionResponse) GetData() MsgVpnMqttSession`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnMqttSessionResponse) GetDataOk() (*MsgVpnMqttSession, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnMqttSessionResponse) SetData(v MsgVpnMqttSession)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnMqttSessionResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnMqttSessionResponse) GetLinks() MsgVpnMqttSessionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnMqttSessionResponse) GetLinksOk() (*MsgVpnMqttSessionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnMqttSessionResponse) SetLinks(v MsgVpnMqttSessionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnMqttSessionResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnMqttSessionResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnMqttSessionResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnMqttSessionResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


