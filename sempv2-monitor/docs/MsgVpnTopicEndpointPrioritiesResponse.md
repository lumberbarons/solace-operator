# MsgVpnTopicEndpointPrioritiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**[]MsgVpnTopicEndpointPriority**](MsgVpnTopicEndpointPriority.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnTopicEndpointPriorityLinks**](MsgVpnTopicEndpointPriorityLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnTopicEndpointPrioritiesResponse

`func NewMsgVpnTopicEndpointPrioritiesResponse(meta SempMeta, ) *MsgVpnTopicEndpointPrioritiesResponse`

NewMsgVpnTopicEndpointPrioritiesResponse instantiates a new MsgVpnTopicEndpointPrioritiesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnTopicEndpointPrioritiesResponseWithDefaults

`func NewMsgVpnTopicEndpointPrioritiesResponseWithDefaults() *MsgVpnTopicEndpointPrioritiesResponse`

NewMsgVpnTopicEndpointPrioritiesResponseWithDefaults instantiates a new MsgVpnTopicEndpointPrioritiesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnTopicEndpointPrioritiesResponse) GetCollections() []map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnTopicEndpointPrioritiesResponse) GetCollectionsOk() (*[]map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnTopicEndpointPrioritiesResponse) SetCollections(v []map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnTopicEndpointPrioritiesResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnTopicEndpointPrioritiesResponse) GetData() []MsgVpnTopicEndpointPriority`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnTopicEndpointPrioritiesResponse) GetDataOk() (*[]MsgVpnTopicEndpointPriority, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnTopicEndpointPrioritiesResponse) SetData(v []MsgVpnTopicEndpointPriority)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnTopicEndpointPrioritiesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnTopicEndpointPrioritiesResponse) GetLinks() []MsgVpnTopicEndpointPriorityLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnTopicEndpointPrioritiesResponse) GetLinksOk() (*[]MsgVpnTopicEndpointPriorityLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnTopicEndpointPrioritiesResponse) SetLinks(v []MsgVpnTopicEndpointPriorityLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnTopicEndpointPrioritiesResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnTopicEndpointPrioritiesResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnTopicEndpointPrioritiesResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnTopicEndpointPrioritiesResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


