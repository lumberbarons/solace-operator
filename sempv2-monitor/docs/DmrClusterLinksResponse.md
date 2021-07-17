# DmrClusterLinksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to [**[]DmrClusterLinkCollections**](DmrClusterLinkCollections.md) |  | [optional] 
**Data** | Pointer to [**[]DmrClusterLink**](DmrClusterLink.md) |  | [optional] 
**Links** | Pointer to [**[]DmrClusterLinkLinks**](DmrClusterLinkLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewDmrClusterLinksResponse

`func NewDmrClusterLinksResponse(meta SempMeta, ) *DmrClusterLinksResponse`

NewDmrClusterLinksResponse instantiates a new DmrClusterLinksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDmrClusterLinksResponseWithDefaults

`func NewDmrClusterLinksResponseWithDefaults() *DmrClusterLinksResponse`

NewDmrClusterLinksResponseWithDefaults instantiates a new DmrClusterLinksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *DmrClusterLinksResponse) GetCollections() []DmrClusterLinkCollections`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *DmrClusterLinksResponse) GetCollectionsOk() (*[]DmrClusterLinkCollections, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *DmrClusterLinksResponse) SetCollections(v []DmrClusterLinkCollections)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *DmrClusterLinksResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *DmrClusterLinksResponse) GetData() []DmrClusterLink`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DmrClusterLinksResponse) GetDataOk() (*[]DmrClusterLink, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DmrClusterLinksResponse) SetData(v []DmrClusterLink)`

SetData sets Data field to given value.

### HasData

`func (o *DmrClusterLinksResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *DmrClusterLinksResponse) GetLinks() []DmrClusterLinkLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DmrClusterLinksResponse) GetLinksOk() (*[]DmrClusterLinkLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DmrClusterLinksResponse) SetLinks(v []DmrClusterLinkLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DmrClusterLinksResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *DmrClusterLinksResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DmrClusterLinksResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DmrClusterLinksResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


