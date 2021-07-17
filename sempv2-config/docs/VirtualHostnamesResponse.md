# VirtualHostnamesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]VirtualHostname**](VirtualHostname.md) |  | [optional] 
**Links** | Pointer to [**[]VirtualHostnameLinks**](VirtualHostnameLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewVirtualHostnamesResponse

`func NewVirtualHostnamesResponse(meta SempMeta, ) *VirtualHostnamesResponse`

NewVirtualHostnamesResponse instantiates a new VirtualHostnamesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualHostnamesResponseWithDefaults

`func NewVirtualHostnamesResponseWithDefaults() *VirtualHostnamesResponse`

NewVirtualHostnamesResponseWithDefaults instantiates a new VirtualHostnamesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *VirtualHostnamesResponse) GetData() []VirtualHostname`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VirtualHostnamesResponse) GetDataOk() (*[]VirtualHostname, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VirtualHostnamesResponse) SetData(v []VirtualHostname)`

SetData sets Data field to given value.

### HasData

`func (o *VirtualHostnamesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *VirtualHostnamesResponse) GetLinks() []VirtualHostnameLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *VirtualHostnamesResponse) GetLinksOk() (*[]VirtualHostnameLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *VirtualHostnamesResponse) SetLinks(v []VirtualHostnameLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *VirtualHostnamesResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *VirtualHostnamesResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *VirtualHostnamesResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *VirtualHostnamesResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


