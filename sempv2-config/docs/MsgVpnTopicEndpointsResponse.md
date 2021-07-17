# MsgVpnTopicEndpointsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]MsgVpnTopicEndpoint**](MsgVpnTopicEndpoint.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnTopicEndpointLinks**](MsgVpnTopicEndpointLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnTopicEndpointsResponse

`func NewMsgVpnTopicEndpointsResponse(meta SempMeta, ) *MsgVpnTopicEndpointsResponse`

NewMsgVpnTopicEndpointsResponse instantiates a new MsgVpnTopicEndpointsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnTopicEndpointsResponseWithDefaults

`func NewMsgVpnTopicEndpointsResponseWithDefaults() *MsgVpnTopicEndpointsResponse`

NewMsgVpnTopicEndpointsResponseWithDefaults instantiates a new MsgVpnTopicEndpointsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MsgVpnTopicEndpointsResponse) GetData() []MsgVpnTopicEndpoint`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnTopicEndpointsResponse) GetDataOk() (*[]MsgVpnTopicEndpoint, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnTopicEndpointsResponse) SetData(v []MsgVpnTopicEndpoint)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnTopicEndpointsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnTopicEndpointsResponse) GetLinks() []MsgVpnTopicEndpointLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnTopicEndpointsResponse) GetLinksOk() (*[]MsgVpnTopicEndpointLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnTopicEndpointsResponse) SetLinks(v []MsgVpnTopicEndpointLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnTopicEndpointsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnTopicEndpointsResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnTopicEndpointsResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnTopicEndpointsResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


