# VirtualHostnameResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to **map[string]interface{}** |  | [optional] 
**Data** | Pointer to [**VirtualHostname**](VirtualHostname.md) |  | [optional] 
**Links** | Pointer to [**VirtualHostnameLinks**](VirtualHostnameLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewVirtualHostnameResponse

`func NewVirtualHostnameResponse(meta SempMeta, ) *VirtualHostnameResponse`

NewVirtualHostnameResponse instantiates a new VirtualHostnameResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualHostnameResponseWithDefaults

`func NewVirtualHostnameResponseWithDefaults() *VirtualHostnameResponse`

NewVirtualHostnameResponseWithDefaults instantiates a new VirtualHostnameResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *VirtualHostnameResponse) GetCollections() map[string]interface{}`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *VirtualHostnameResponse) GetCollectionsOk() (*map[string]interface{}, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *VirtualHostnameResponse) SetCollections(v map[string]interface{})`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *VirtualHostnameResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *VirtualHostnameResponse) GetData() VirtualHostname`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VirtualHostnameResponse) GetDataOk() (*VirtualHostname, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VirtualHostnameResponse) SetData(v VirtualHostname)`

SetData sets Data field to given value.

### HasData

`func (o *VirtualHostnameResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *VirtualHostnameResponse) GetLinks() VirtualHostnameLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VirtualHostnameResponse) GetLinksOk() (*VirtualHostnameLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VirtualHostnameResponse) SetLinks(v VirtualHostnameLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *VirtualHostnameResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *VirtualHostnameResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *VirtualHostnameResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *VirtualHostnameResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


