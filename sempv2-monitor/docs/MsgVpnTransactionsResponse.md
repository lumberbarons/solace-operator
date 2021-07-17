# MsgVpnTransactionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to [**[]MsgVpnTransactionCollections**](MsgVpnTransactionCollections.md) |  | [optional] 
**Data** | Pointer to [**[]MsgVpnTransaction**](MsgVpnTransaction.md) |  | [optional] 
**Links** | Pointer to [**[]MsgVpnTransactionLinks**](MsgVpnTransactionLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewMsgVpnTransactionsResponse

`func NewMsgVpnTransactionsResponse(meta SempMeta, ) *MsgVpnTransactionsResponse`

NewMsgVpnTransactionsResponse instantiates a new MsgVpnTransactionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsgVpnTransactionsResponseWithDefaults

`func NewMsgVpnTransactionsResponseWithDefaults() *MsgVpnTransactionsResponse`

NewMsgVpnTransactionsResponseWithDefaults instantiates a new MsgVpnTransactionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *MsgVpnTransactionsResponse) GetCollections() []MsgVpnTransactionCollections`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *MsgVpnTransactionsResponse) GetCollectionsOk() (*[]MsgVpnTransactionCollections, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *MsgVpnTransactionsResponse) SetCollections(v []MsgVpnTransactionCollections)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *MsgVpnTransactionsResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *MsgVpnTransactionsResponse) GetData() []MsgVpnTransaction`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MsgVpnTransactionsResponse) GetDataOk() (*[]MsgVpnTransaction, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MsgVpnTransactionsResponse) SetData(v []MsgVpnTransaction)`

SetData sets Data field to given value.

### HasData

`func (o *MsgVpnTransactionsResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *MsgVpnTransactionsResponse) GetLinks() []MsgVpnTransactionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MsgVpnTransactionsResponse) GetLinksOk() (*[]MsgVpnTransactionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MsgVpnTransactionsResponse) SetLinks(v []MsgVpnTransactionLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MsgVpnTransactionsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *MsgVpnTransactionsResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *MsgVpnTransactionsResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *MsgVpnTransactionsResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


