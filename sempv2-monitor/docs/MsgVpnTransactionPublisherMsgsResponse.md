# MsgVpnTransactionPublisherMsgsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**[]MsgVpnTransactionPublisherMsg**](MsgVpnTransactionPublisherMsg.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnTransactionPublisherMsgLinks**](MsgVpnTransactionPublisherMsgLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnTransactionPublisherMsgsResponse

`func NewMsgVpnTransactionPublisherMsgsResponse(meta SempMeta, ) *MsgVpnTransactionPublisherMsgsResponse`

NewMsgVpnTransactionPublisherMsgsResponse instantiates a new MsgVpnTransactionPublisherMsgsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnTransactionPublisherMsgsResponseWithDefaults

`func NewMsgVpnTransactionPublisherMsgsResponseWithDefaults() *MsgVpnTransactionPublisherMsgsResponse`

NewMsgVpnTransactionPublisherMsgsResponseWithDefaults instantiates a new MsgVpnTransactionPublisherMsgsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnTransactionPublisherMsgsResponse) GetCollections() []map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnTransactionPublisherMsgsResponse) GetCollectionsOk() (*[]map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnTransactionPublisherMsgsResponse) SetCollections(v []map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnTransactionPublisherMsgsResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnTransactionPublisherMsgsResponse) GetData() []MsgVpnTransactionPublisherMsg`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnTransactionPublisherMsgsResponse) GetDataOk() (*[]MsgVpnTransactionPublisherMsg, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnTransactionPublisherMsgsResponse) SetData(v []MsgVpnTransactionPublisherMsg)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnTransactionPublisherMsgsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnTransactionPublisherMsgsResponse) GetLinks() []MsgVpnTransactionPublisherMsgLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnTransactionPublisherMsgsResponse) GetLinksOk() (*[]MsgVpnTransactionPublisherMsgLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnTransactionPublisherMsgsResponse) SetLinks(v []MsgVpnTransactionPublisherMsgLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnTransactionPublisherMsgsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnTransactionPublisherMsgsResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnTransactionPublisherMsgsResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnTransactionPublisherMsgsResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


