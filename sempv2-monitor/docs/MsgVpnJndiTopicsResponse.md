# MsgVpnJndiTopicsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**[]MsgVpnJndiTopic**](MsgVpnJndiTopic.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnJndiTopicLinks**](MsgVpnJndiTopicLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnJndiTopicsResponse

`func NewMsgVpnJndiTopicsResponse(meta SempMeta, ) *MsgVpnJndiTopicsResponse`

NewMsgVpnJndiTopicsResponse instantiates a new MsgVpnJndiTopicsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnJndiTopicsResponseWithDefaults

`func NewMsgVpnJndiTopicsResponseWithDefaults() *MsgVpnJndiTopicsResponse`

NewMsgVpnJndiTopicsResponseWithDefaults instantiates a new MsgVpnJndiTopicsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnJndiTopicsResponse) GetCollections() []map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnJndiTopicsResponse) GetCollectionsOk() (*[]map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnJndiTopicsResponse) SetCollections(v []map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnJndiTopicsResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnJndiTopicsResponse) GetData() []MsgVpnJndiTopic`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnJndiTopicsResponse) GetDataOk() (*[]MsgVpnJndiTopic, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnJndiTopicsResponse) SetData(v []MsgVpnJndiTopic)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnJndiTopicsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnJndiTopicsResponse) GetLinks() []MsgVpnJndiTopicLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnJndiTopicsResponse) GetLinksOk() (*[]MsgVpnJndiTopicLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnJndiTopicsResponse) SetLinks(v []MsgVpnJndiTopicLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnJndiTopicsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnJndiTopicsResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnJndiTopicsResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnJndiTopicsResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


