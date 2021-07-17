# MsgVpnReplicatedTopicResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**MsgVpnReplicatedTopic**](MsgVpnReplicatedTopic.md) |  | [optional] 
**Links** | Pointer to [**MsgVpnReplicatedTopicLinks**](MsgVpnReplicatedTopicLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnReplicatedTopicResponse

`func NewMsgVpnReplicatedTopicResponse(meta SempMeta, ) *MsgVpnReplicatedTopicResponse`

NewMsgVpnReplicatedTopicResponse instantiates a new MsgVpnReplicatedTopicResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnReplicatedTopicResponseWithDefaults

`func NewMsgVpnReplicatedTopicResponseWithDefaults() *MsgVpnReplicatedTopicResponse`

NewMsgVpnReplicatedTopicResponseWithDefaults instantiates a new MsgVpnReplicatedTopicResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnReplicatedTopicResponse) GetCollections() map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnReplicatedTopicResponse) GetCollectionsOk() (*map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnReplicatedTopicResponse) SetCollections(v map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnReplicatedTopicResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnReplicatedTopicResponse) GetData() MsgVpnReplicatedTopic`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnReplicatedTopicResponse) GetDataOk() (*MsgVpnReplicatedTopic, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnReplicatedTopicResponse) SetData(v MsgVpnReplicatedTopic)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnReplicatedTopicResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnReplicatedTopicResponse) GetLinks() MsgVpnReplicatedTopicLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnReplicatedTopicResponse) GetLinksOk() (*MsgVpnReplicatedTopicLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnReplicatedTopicResponse) SetLinks(v MsgVpnReplicatedTopicLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnReplicatedTopicResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnReplicatedTopicResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnReplicatedTopicResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnReplicatedTopicResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


