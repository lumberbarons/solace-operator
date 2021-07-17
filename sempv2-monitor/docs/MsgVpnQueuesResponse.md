# MsgVpnQueuesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to [**[]MsgVpnQueueCollections**](MsgVpnQueueCollections.md) |  | [optional] 
**Data** | Pointer to [**[]MsgVpnQueue**](MsgVpnQueue.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnQueueLinks**](MsgVpnQueueLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnQueuesResponse

`func NewMsgVpnQueuesResponse(meta SempMeta, ) *MsgVpnQueuesResponse`

NewMsgVpnQueuesResponse instantiates a new MsgVpnQueuesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnQueuesResponseWithDefaults

`func NewMsgVpnQueuesResponseWithDefaults() *MsgVpnQueuesResponse`

NewMsgVpnQueuesResponseWithDefaults instantiates a new MsgVpnQueuesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnQueuesResponse) GetCollections() []MsgVpnQueueCollections`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnQueuesResponse) GetCollectionsOk() (*[]MsgVpnQueueCollections, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnQueuesResponse) SetCollections(v []MsgVpnQueueCollections)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnQueuesResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnQueuesResponse) GetData() []MsgVpnQueue`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnQueuesResponse) GetDataOk() (*[]MsgVpnQueue, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnQueuesResponse) SetData(v []MsgVpnQueue)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnQueuesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnQueuesResponse) GetLinks() []MsgVpnQueueLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnQueuesResponse) GetLinksOk() (*[]MsgVpnQueueLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnQueuesResponse) SetLinks(v []MsgVpnQueueLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnQueuesResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnQueuesResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnQueuesResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnQueuesResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


