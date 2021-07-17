# CertAuthoritiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collections** | Pointer to [**[]CertAuthorityCollections**](CertAuthorityCollections.md) |  | [optional] 
**Data** | Pointer to [**[]CertAuthority**](CertAuthority.md) |  | [optional] 
**Links** | Pointer to [**[]CertAuthorityLinks**](CertAuthorityLinks.md) |  | [optional] 
**Meta** | [**SempMeta**](SempMeta.md) |  | 

## Methods

### NewCertAuthoritiesResponse

`func NewCertAuthoritiesResponse(meta SempMeta, ) *CertAuthoritiesResponse`

NewCertAuthoritiesResponse instantiates a new CertAuthoritiesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertAuthoritiesResponseWithDefaults

`func NewCertAuthoritiesResponseWithDefaults() *CertAuthoritiesResponse`

NewCertAuthoritiesResponseWithDefaults instantiates a new CertAuthoritiesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollections

`func (o *CertAuthoritiesResponse) GetCollections() []CertAuthorityCollections`

GetCollections returns the Collections field if non-nil, zero value otherwise.

### GetCollectionsOk

`func (o *CertAuthoritiesResponse) GetCollectionsOk() (*[]CertAuthorityCollections, bool)`

GetCollectionsOk returns a tuple with the Collections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollections

`func (o *CertAuthoritiesResponse) SetCollections(v []CertAuthorityCollections)`

SetCollections sets Collections field to given value.

### HasCollections

`func (o *CertAuthoritiesResponse) HasCollections() bool`

HasCollections returns a boolean if a field has been set.

### GetData

`func (o *CertAuthoritiesResponse) GetData() []CertAuthority`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CertAuthoritiesResponse) GetDataOk() (*[]CertAuthority, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CertAuthoritiesResponse) SetData(v []CertAuthority)`

SetData sets Data field to given value.

### HasData

`func (o *CertAuthoritiesResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *CertAuthoritiesResponse) GetLinks() []CertAuthorityLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CertAuthoritiesResponse) GetLinksOk() (*[]CertAuthorityLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CertAuthoritiesResponse) SetLinks(v []CertAuthorityLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CertAuthoritiesResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *CertAuthoritiesResponse) GetMeta() SempMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CertAuthoritiesResponse) GetMetaOk() (*SempMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CertAuthoritiesResponse) SetMeta(v SempMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


