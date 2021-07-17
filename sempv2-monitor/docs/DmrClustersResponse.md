# DmrClustersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to [**[]DmrClusterCollections**](DmrClusterCollections.md) |  | [optional] 
**Data** | Pointer to [**[]DmrCluster**](DmrCluster.md) |  | [optional] 
**Links** | Pointer to [**[]DmrClusterLinks**](DmrClusterLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewDmrClustersResponse

`func NewDmrClustersResponse(meta SempMeta, ) *DmrClustersResponse`

NewDmrClustersResponse instantiates a new DmrClustersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDmrClustersResponseWithDefaults

`func NewDmrClustersResponseWithDefaults() *DmrClustersResponse`

NewDmrClustersResponseWithDefaults instantiates a new DmrClustersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *DmrClustersResponse) GetCollections() []DmrClusterCollections`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *DmrClustersResponse) GetCollectionsOk() (*[]DmrClusterCollections, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *DmrClustersResponse) SetCollections(v []DmrClusterCollections)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *DmrClustersResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *DmrClustersResponse) GetData() []DmrCluster`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DmrClustersResponse) GetDataOk() (*[]DmrCluster, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DmrClustersResponse) SetData(v []DmrCluster)`

SetData sets Data field to given value.

### HasData

`func (o *DmrClustersResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *DmrClustersResponse) GetLinks() []DmrClusterLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DmrClustersResponse) GetLinksOk() (*[]DmrClusterLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DmrClustersResponse) SetLinks(v []DmrClusterLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DmrClustersResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *DmrClustersResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DmrClustersResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DmrClustersResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


