# MsgVpnMqttSessionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to [**[]MsgVpnMqttSessionCollections**](MsgVpnMqttSessionCollections.md) |  | [optional] 
**Data** | Pointer to [**[]MsgVpnMqttSession**](MsgVpnMqttSession.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnMqttSessionLinks**](MsgVpnMqttSessionLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnMqttSessionsResponse

`func NewMsgVpnMqttSessionsResponse(meta SempMeta, ) *MsgVpnMqttSessionsResponse`

NewMsgVpnMqttSessionsResponse instantiates a new MsgVpnMqttSessionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnMqttSessionsResponseWithDefaults

`func NewMsgVpnMqttSessionsResponseWithDefaults() *MsgVpnMqttSessionsResponse`

NewMsgVpnMqttSessionsResponseWithDefaults instantiates a new MsgVpnMqttSessionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnMqttSessionsResponse) GetCollections() []MsgVpnMqttSessionCollections`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnMqttSessionsResponse) GetCollectionsOk() (*[]MsgVpnMqttSessionCollections, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnMqttSessionsResponse) SetCollections(v []MsgVpnMqttSessionCollections)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnMqttSessionsResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnMqttSessionsResponse) GetData() []MsgVpnMqttSession`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnMqttSessionsResponse) GetDataOk() (*[]MsgVpnMqttSession, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnMqttSessionsResponse) SetData(v []MsgVpnMqttSession)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnMqttSessionsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnMqttSessionsResponse) GetLinks() []MsgVpnMqttSessionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnMqttSessionsResponse) GetLinksOk() (*[]MsgVpnMqttSessionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnMqttSessionsResponse) SetLinks(v []MsgVpnMqttSessionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnMqttSessionsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnMqttSessionsResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnMqttSessionsResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnMqttSessionsResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


