# DmrClusterResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**DmrCluster**](DmrCluster.md) |  | [optional] 
**Links** | Pointer to [**DmrClusterLinks**](DmrClusterLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewDmrClusterResponse

`func NewDmrClusterResponse(meta SempMeta, ) *DmrClusterResponse`

NewDmrClusterResponse instantiates a new DmrClusterResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDmrClusterResponseWithDefaults

`func NewDmrClusterResponseWithDefaults() *DmrClusterResponse`

NewDmrClusterResponseWithDefaults instantiates a new DmrClusterResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DmrClusterResponse) GetData() DmrCluster`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DmrClusterResponse) GetDataOk() (*DmrCluster, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DmrClusterResponse) SetData(v DmrCluster)`

SetData sets Data field to given value.

### HasData

`func (o *DmrClusterResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *DmrClusterResponse) GetLinks() DmrClusterLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DmrClusterResponse) GetLinksOk() (*DmrClusterLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DmrClusterResponse) SetLinks(v DmrClusterLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DmrClusterResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *DmrClusterResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DmrClusterResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DmrClusterResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


