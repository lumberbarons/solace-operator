# MsgVpnJndiQueuesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]MsgVpnJndiQueue**](MsgVpnJndiQueue.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnJndiQueueLinks**](MsgVpnJndiQueueLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnJndiQueuesResponse

`func NewMsgVpnJndiQueuesResponse(meta SempMeta, ) *MsgVpnJndiQueuesResponse`

NewMsgVpnJndiQueuesResponse instantiates a new MsgVpnJndiQueuesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnJndiQueuesResponseWithDefaults

`func NewMsgVpnJndiQueuesResponseWithDefaults() *MsgVpnJndiQueuesResponse`

NewMsgVpnJndiQueuesResponseWithDefaults instantiates a new MsgVpnJndiQueuesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MsgVpnJndiQueuesResponse) GetData() []MsgVpnJndiQueue`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnJndiQueuesResponse) GetDataOk() (*[]MsgVpnJndiQueue, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnJndiQueuesResponse) SetData(v []MsgVpnJndiQueue)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnJndiQueuesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnJndiQueuesResponse) GetLinks() []MsgVpnJndiQueueLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnJndiQueuesResponse) GetLinksOk() (*[]MsgVpnJndiQueueLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnJndiQueuesResponse) SetLinks(v []MsgVpnJndiQueueLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnJndiQueuesResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnJndiQueuesResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnJndiQueuesResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnJndiQueuesResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


