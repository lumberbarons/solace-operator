# MsgVpnSequencedTopicResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**MsgVpnSequencedTopic**](MsgVpnSequencedTopic.md) |  | [optional] 
**Links** | Pointer to [**MsgVpnSequencedTopicLinks**](MsgVpnSequencedTopicLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnSequencedTopicResponse

`func NewMsgVpnSequencedTopicResponse(meta SempMeta, ) *MsgVpnSequencedTopicResponse`

NewMsgVpnSequencedTopicResponse instantiates a new MsgVpnSequencedTopicResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnSequencedTopicResponseWithDefaults

`func NewMsgVpnSequencedTopicResponseWithDefaults() *MsgVpnSequencedTopicResponse`

NewMsgVpnSequencedTopicResponseWithDefaults instantiates a new MsgVpnSequencedTopicResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MsgVpnSequencedTopicResponse) GetData() MsgVpnSequencedTopic`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnSequencedTopicResponse) GetDataOk() (*MsgVpnSequencedTopic, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnSequencedTopicResponse) SetData(v MsgVpnSequencedTopic)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnSequencedTopicResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnSequencedTopicResponse) GetLinks() MsgVpnSequencedTopicLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnSequencedTopicResponse) GetLinksOk() (*MsgVpnSequencedTopicLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnSequencedTopicResponse) SetLinks(v MsgVpnSequencedTopicLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnSequencedTopicResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnSequencedTopicResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnSequencedTopicResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnSequencedTopicResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


