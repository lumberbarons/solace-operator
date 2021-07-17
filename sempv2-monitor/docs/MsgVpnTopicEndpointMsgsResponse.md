# MsgVpnTopicEndpointMsgsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**[]MsgVpnTopicEndpointMsg**](MsgVpnTopicEndpointMsg.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnTopicEndpointMsgLinks**](MsgVpnTopicEndpointMsgLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnTopicEndpointMsgsResponse

`func NewMsgVpnTopicEndpointMsgsResponse(meta SempMeta, ) *MsgVpnTopicEndpointMsgsResponse`

NewMsgVpnTopicEndpointMsgsResponse instantiates a new MsgVpnTopicEndpointMsgsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnTopicEndpointMsgsResponseWithDefaults

`func NewMsgVpnTopicEndpointMsgsResponseWithDefaults() *MsgVpnTopicEndpointMsgsResponse`

NewMsgVpnTopicEndpointMsgsResponseWithDefaults instantiates a new MsgVpnTopicEndpointMsgsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnTopicEndpointMsgsResponse) GetCollections() []map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnTopicEndpointMsgsResponse) GetCollectionsOk() (*[]map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnTopicEndpointMsgsResponse) SetCollections(v []map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnTopicEndpointMsgsResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnTopicEndpointMsgsResponse) GetData() []MsgVpnTopicEndpointMsg`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnTopicEndpointMsgsResponse) GetDataOk() (*[]MsgVpnTopicEndpointMsg, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnTopicEndpointMsgsResponse) SetData(v []MsgVpnTopicEndpointMsg)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnTopicEndpointMsgsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnTopicEndpointMsgsResponse) GetLinks() []MsgVpnTopicEndpointMsgLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnTopicEndpointMsgsResponse) GetLinksOk() (*[]MsgVpnTopicEndpointMsgLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnTopicEndpointMsgsResponse) SetLinks(v []MsgVpnTopicEndpointMsgLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnTopicEndpointMsgsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnTopicEndpointMsgsResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnTopicEndpointMsgsResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnTopicEndpointMsgsResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


