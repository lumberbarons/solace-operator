# MsgVpnReplicatedTopicsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]MsgVpnReplicatedTopic**](MsgVpnReplicatedTopic.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnReplicatedTopicLinks**](MsgVpnReplicatedTopicLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnReplicatedTopicsResponse

`func NewMsgVpnReplicatedTopicsResponse(meta SempMeta, ) *MsgVpnReplicatedTopicsResponse`

NewMsgVpnReplicatedTopicsResponse instantiates a new MsgVpnReplicatedTopicsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnReplicatedTopicsResponseWithDefaults

`func NewMsgVpnReplicatedTopicsResponseWithDefaults() *MsgVpnReplicatedTopicsResponse`

NewMsgVpnReplicatedTopicsResponseWithDefaults instantiates a new MsgVpnReplicatedTopicsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MsgVpnReplicatedTopicsResponse) GetData() []MsgVpnReplicatedTopic`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnReplicatedTopicsResponse) GetDataOk() (*[]MsgVpnReplicatedTopic, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnReplicatedTopicsResponse) SetData(v []MsgVpnReplicatedTopic)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnReplicatedTopicsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnReplicatedTopicsResponse) GetLinks() []MsgVpnReplicatedTopicLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnReplicatedTopicsResponse) GetLinksOk() (*[]MsgVpnReplicatedTopicLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnReplicatedTopicsResponse) SetLinks(v []MsgVpnReplicatedTopicLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnReplicatedTopicsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnReplicatedTopicsResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnReplicatedTopicsResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnReplicatedTopicsResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


