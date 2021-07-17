# BrokerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**Broker**](Broker.md) |  | [optional] 
**Links** | Pointer to [**BrokerLinks**](BrokerLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewBrokerResponse

`func NewBrokerResponse(meta SempMeta, ) *BrokerResponse`

NewBrokerResponse instantiates a new BrokerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrokerResponseWithDefaults

`func NewBrokerResponseWithDefaults() *BrokerResponse`

NewBrokerResponseWithDefaults instantiates a new BrokerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BrokerResponse) GetData() Broker`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BrokerResponse) GetDataOk() (*Broker, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BrokerResponse) SetData(v Broker)`

SetData sets Data field to given value.

### HasData

`func (o *BrokerResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *BrokerResponse) GetLinks() BrokerLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BrokerResponse) GetLinksOk() (*BrokerLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BrokerResponse) SetLinks(v BrokerLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BrokerResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *BrokerResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BrokerResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BrokerResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


