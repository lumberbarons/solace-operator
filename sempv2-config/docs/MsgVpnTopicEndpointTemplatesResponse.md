# MsgVpnTopicEndpointTemplatesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]MsgVpnTopicEndpointTemplate**](MsgVpnTopicEndpointTemplate.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnTopicEndpointTemplateLinks**](MsgVpnTopicEndpointTemplateLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnTopicEndpointTemplatesResponse

`func NewMsgVpnTopicEndpointTemplatesResponse(meta SempMeta, ) *MsgVpnTopicEndpointTemplatesResponse`

NewMsgVpnTopicEndpointTemplatesResponse instantiates a new MsgVpnTopicEndpointTemplatesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnTopicEndpointTemplatesResponseWithDefaults

`func NewMsgVpnTopicEndpointTemplatesResponseWithDefaults() *MsgVpnTopicEndpointTemplatesResponse`

NewMsgVpnTopicEndpointTemplatesResponseWithDefaults instantiates a new MsgVpnTopicEndpointTemplatesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MsgVpnTopicEndpointTemplatesResponse) GetData() []MsgVpnTopicEndpointTemplate`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnTopicEndpointTemplatesResponse) GetDataOk() (*[]MsgVpnTopicEndpointTemplate, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnTopicEndpointTemplatesResponse) SetData(v []MsgVpnTopicEndpointTemplate)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnTopicEndpointTemplatesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnTopicEndpointTemplatesResponse) GetLinks() []MsgVpnTopicEndpointTemplateLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnTopicEndpointTemplatesResponse) GetLinksOk() (*[]MsgVpnTopicEndpointTemplateLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnTopicEndpointTemplatesResponse) SetLinks(v []MsgVpnTopicEndpointTemplateLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnTopicEndpointTemplatesResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnTopicEndpointTemplatesResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnTopicEndpointTemplatesResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnTopicEndpointTemplatesResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


