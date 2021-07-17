# DmrClusterLinkRemoteAddressResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**DmrClusterLinkRemoteAddress**](DmrClusterLinkRemoteAddress.md) |  | [optional] 
**Links** | Pointer to [**DmrClusterLinkRemoteAddressLinks**](DmrClusterLinkRemoteAddressLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewDmrClusterLinkRemoteAddressResponse

`func NewDmrClusterLinkRemoteAddressResponse(meta SempMeta, ) *DmrClusterLinkRemoteAddressResponse`

NewDmrClusterLinkRemoteAddressResponse instantiates a new DmrClusterLinkRemoteAddressResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDmrClusterLinkRemoteAddressResponseWithDefaults

`func NewDmrClusterLinkRemoteAddressResponseWithDefaults() *DmrClusterLinkRemoteAddressResponse`

NewDmrClusterLinkRemoteAddressResponseWithDefaults instantiates a new DmrClusterLinkRemoteAddressResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DmrClusterLinkRemoteAddressResponse) GetData() DmrClusterLinkRemoteAddress`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DmrClusterLinkRemoteAddressResponse) GetDataOk() (*DmrClusterLinkRemoteAddress, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DmrClusterLinkRemoteAddressResponse) SetData(v DmrClusterLinkRemoteAddress)`

SetData sets Data field to given value.

### HasData

`func (o *DmrClusterLinkRemoteAddressResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *DmrClusterLinkRemoteAddressResponse) GetLinks() DmrClusterLinkRemoteAddressLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DmrClusterLinkRemoteAddressResponse) GetLinksOk() (*DmrClusterLinkRemoteAddressLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DmrClusterLinkRemoteAddressResponse) SetLinks(v DmrClusterLinkRemoteAddressLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DmrClusterLinkRemoteAddressResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *DmrClusterLinkRemoteAddressResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DmrClusterLinkRemoteAddressResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DmrClusterLinkRemoteAddressResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


