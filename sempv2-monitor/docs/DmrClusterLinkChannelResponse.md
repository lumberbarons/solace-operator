# DmrClusterLinkChannelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**DmrClusterLinkChannel**](DmrClusterLinkChannel.md) |  | [optional] 
**Links** | Pointer to [**DmrClusterLinkChannelLinks**](DmrClusterLinkChannelLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewDmrClusterLinkChannelResponse

`func NewDmrClusterLinkChannelResponse(meta SempMeta, ) *DmrClusterLinkChannelResponse`

NewDmrClusterLinkChannelResponse instantiates a new DmrClusterLinkChannelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDmrClusterLinkChannelResponseWithDefaults

`func NewDmrClusterLinkChannelResponseWithDefaults() *DmrClusterLinkChannelResponse`

NewDmrClusterLinkChannelResponseWithDefaults instantiates a new DmrClusterLinkChannelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *DmrClusterLinkChannelResponse) GetCollections() map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *DmrClusterLinkChannelResponse) GetCollectionsOk() (*map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *DmrClusterLinkChannelResponse) SetCollections(v map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *DmrClusterLinkChannelResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *DmrClusterLinkChannelResponse) GetData() DmrClusterLinkChannel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DmrClusterLinkChannelResponse) GetDataOk() (*DmrClusterLinkChannel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DmrClusterLinkChannelResponse) SetData(v DmrClusterLinkChannel)`

SetData sets Data field to given value.

### HasData

`func (o *DmrClusterLinkChannelResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *DmrClusterLinkChannelResponse) GetLinks() DmrClusterLinkChannelLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DmrClusterLinkChannelResponse) GetLinksOk() (*DmrClusterLinkChannelLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DmrClusterLinkChannelResponse) SetLinks(v DmrClusterLinkChannelLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DmrClusterLinkChannelResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *DmrClusterLinkChannelResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DmrClusterLinkChannelResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DmrClusterLinkChannelResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


