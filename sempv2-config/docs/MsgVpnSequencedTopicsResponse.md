# MsgVpnSequencedTopicsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]MsgVpnSequencedTopic**](MsgVpnSequencedTopic.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnSequencedTopicLinks**](MsgVpnSequencedTopicLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnSequencedTopicsResponse

`func NewMsgVpnSequencedTopicsResponse(meta SempMeta, ) *MsgVpnSequencedTopicsResponse`

NewMsgVpnSequencedTopicsResponse instantiates a new MsgVpnSequencedTopicsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnSequencedTopicsResponseWithDefaults

`func NewMsgVpnSequencedTopicsResponseWithDefaults() *MsgVpnSequencedTopicsResponse`

NewMsgVpnSequencedTopicsResponseWithDefaults instantiates a new MsgVpnSequencedTopicsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MsgVpnSequencedTopicsResponse) GetData() []MsgVpnSequencedTopic`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnSequencedTopicsResponse) GetDataOk() (*[]MsgVpnSequencedTopic, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnSequencedTopicsResponse) SetData(v []MsgVpnSequencedTopic)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnSequencedTopicsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnSequencedTopicsResponse) GetLinks() []MsgVpnSequencedTopicLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnSequencedTopicsResponse) GetLinksOk() (*[]MsgVpnSequencedTopicLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnSequencedTopicsResponse) SetLinks(v []MsgVpnSequencedTopicLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnSequencedTopicsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnSequencedTopicsResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnSequencedTopicsResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnSequencedTopicsResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


