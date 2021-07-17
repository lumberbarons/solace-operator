# DmrClusterLinkResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**DmrClusterLink**](DmrClusterLink.md) |  | [optional] 
**Links** | Pointer to [**DmrClusterLinkLinks**](DmrClusterLinkLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewDmrClusterLinkResponse

`func NewDmrClusterLinkResponse(meta SempMeta, ) *DmrClusterLinkResponse`

NewDmrClusterLinkResponse instantiates a new DmrClusterLinkResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDmrClusterLinkResponseWithDefaults

`func NewDmrClusterLinkResponseWithDefaults() *DmrClusterLinkResponse`

NewDmrClusterLinkResponseWithDefaults instantiates a new DmrClusterLinkResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DmrClusterLinkResponse) GetData() DmrClusterLink`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DmrClusterLinkResponse) GetDataOk() (*DmrClusterLink, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DmrClusterLinkResponse) SetData(v DmrClusterLink)`

SetData sets Data field to given value.

### HasData

`func (o *DmrClusterLinkResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *DmrClusterLinkResponse) GetLinks() DmrClusterLinkLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DmrClusterLinkResponse) GetLinksOk() (*DmrClusterLinkLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DmrClusterLinkResponse) SetLinks(v DmrClusterLinkLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DmrClusterLinkResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *DmrClusterLinkResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DmrClusterLinkResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DmrClusterLinkResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


