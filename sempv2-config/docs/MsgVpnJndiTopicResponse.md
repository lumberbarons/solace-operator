# MsgVpnJndiTopicResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**MsgVpnJndiTopic**](MsgVpnJndiTopic.md) |  | [optional] 
**Links** | Pointer to [**MsgVpnJndiTopicLinks**](MsgVpnJndiTopicLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnJndiTopicResponse

`func NewMsgVpnJndiTopicResponse(meta SempMeta, ) *MsgVpnJndiTopicResponse`

NewMsgVpnJndiTopicResponse instantiates a new MsgVpnJndiTopicResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnJndiTopicResponseWithDefaults

`func NewMsgVpnJndiTopicResponseWithDefaults() *MsgVpnJndiTopicResponse`

NewMsgVpnJndiTopicResponseWithDefaults instantiates a new MsgVpnJndiTopicResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MsgVpnJndiTopicResponse) GetData() MsgVpnJndiTopic`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnJndiTopicResponse) GetDataOk() (*MsgVpnJndiTopic, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnJndiTopicResponse) SetData(v MsgVpnJndiTopic)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnJndiTopicResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnJndiTopicResponse) GetLinks() MsgVpnJndiTopicLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnJndiTopicResponse) GetLinksOk() (*MsgVpnJndiTopicLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnJndiTopicResponse) SetLinks(v MsgVpnJndiTopicLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnJndiTopicResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnJndiTopicResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnJndiTopicResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnJndiTopicResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


